package servers

import (
	"fmt"
	"regexp"
	"strings"

	"ecloud-sdk-go/pkg/errors"
	"ecloud-sdk-go/pkg/types"
)

type ActionType string
type ServersType string
type VMType string
type VolumeType string

const (
	SERVER_TYPE_VM             ServersType = "VM"
	SERVER_TYPE_IRONIC         ServersType = "IRONIC"
	SERVER_TYPE_DECLOUD_SERVER ServersType = "DECLOUD_SERVER"
	VM_TYPE_HIGHIO             VMType      = "highIO"
	VM_TYPE_EXCLUSIVE          VMType      = "exclusive"
	VM_TYPE_MEMIMPROVE         VMType      = "memImprove"
	VM_TYPE_COMMON             VMType      = "common"
	VM_TYPE_GPU                VMType      = "gpu"
	ACTION_PAUSE               ActionType  = "PAUSE"
	ACTION_UNPAUSE             ActionType  = "UNPAUSE"
	ACTION_STOP                ActionType  = "STOP"
	ACTION_START               ActionType  = "START"
	ACTION_LOCK                ActionType  = "LOCK"
	ACTION_UNLOCK              ActionType  = "UNLOCK"
	ACTION_SUSPEND             ActionType  = "SUSPEND"
	ACTION_RESUME              ActionType  = "RESUME"
	ACTION_RESCUE              ActionType  = "RESCUE"
	ACTION_UNRESCUE            ActionType  = "UNRESCUE "
	ACTION_SHELVE              ActionType  = "SHELVE"
	ACTION_SHELVE_OFFLOAD      ActionType  = "SHELVE_OFFLOAD"
	ACTION_UNSHELVE            ActionType  = "UNSHELVE"

	VolumeType_CAPACITY                VolumeType = "capacity"
	VolumeType_PERFORMANCEOPTIMIZATION VolumeType = "performanceOptimization"
	VolumeType_HIGHPERFORMANCE         VolumeType = "highPerformance"
	VolumeType_LOACL                   VolumeType = "local"
)

//different from results of Get(),or List()
//todo: convert results of Get() or List() to CreateOpts
type CreateOpts struct {
	Region           string `json:"region,omitempty"`
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	Name             string `json:"name" required:"true"`
	//get cpu and ram with flavorRef
	Cpu         int    `json:"cpu" required:"true"`
	Ram         int    `json:"ram" required:"true"`
	Disk        int    `json:"disk" required:"true"`
	Gpu         string `json:"gpu,omitempty"`
	ImageName   string `json:"imageName" required:"true"`
	Password    string `json:"password,omitempty" or:"KeypairName"`
	KeypairName string `json:"keypairName,omitempty" or:"Password"`
	//only need set name, size ,type
	BootVolume *Volume `json:"bootVolume,omitempty"`
	//Volumes          []Volume          `json:"volumes"`
	Networks         []ServerNetwork   `json:"networks" required:"true"`
	SecurityGroupIds []string          `json:"securityGroupIds" required:"true"`
	UserData         string            `json:"userData,omitempty"`
	ProductType      types.ProductType `json:"productType" required:"true"`
	ServerType       ServersType       `json:"serverType" required:"true"`
	VMType           VMType            `json:"vmType" required:"true"`
	NeedConfirm      *bool             `json:"needConfirm" required:"true"`
	NeedEncryption   *bool
}

func (o *CreateOpts) Verify() error {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9-]{5,26}$", o.Name)
	if !ok {
		return errors.BaseError{Info: "Name of Server should match regexp ^[a-zA-Z0-9-]{5,26}$"}
	}

	if o.NeedEncryption == nil || *o.NeedEncryption == false {
		return nil
	}

	if o.Password == "" {
		return nil
	}
	numb := 0
	letterUpper := 0
	letterLower := 0
	symbol := 0
	symStrs := "~@#$%*-+=:,.?[]{}_"
	for _, i := range o.Password {
		if i < 58 && i > 47 {
			numb++
			continue
		}
		if i > 64 && i < 91 {
			letterUpper++
			continue
		}
		if i > 96 && i < 123 {
			letterLower++
			continue
		}
		if strings.Contains(symStrs, string(i)) {
			symbol++
			continue
		}
	}
	length := len(o.Password)
	if numb != 0 && letterLower != 0 && letterUpper != 0 && symbol != 0 && symbol <= 3 && length >= 8 && length <= 16 {
		return nil
	} else {
		return errors.BaseError{Info: "密码必须符合：8-16位字符，同时包括数字、大小写字母和特殊字符，其中特殊字符最多不能超过3个，且需要在“~ @ # $ % * _ - + = : , . ? [ ] { }”范围内"}
	}
}

type Volume struct {
	Size       int        `json:"size" required:"true"`
	VolumeType VolumeType `json:"volumeType" required:"true"`
}

//different from network object
type ServerNetwork struct {
	PortID    string `json:"portId,omitempty"`
	NetworkID string `json:"networkId,omitempty"`
	FixedIP   string `json:"fixedIp,omitempty"`
}

//TODO-SML  这里tag q 后续用于拼接query参数，详见
type ListOpts struct {
	Name         string         `q:"serverName"`
	ProductTypes string         `q:"productTypes"`
	ServerId     string         `q:"serverId"`
	OPStatus     OPStatusType   `q:"opStatus"`
	ECStatus     []ECStatusType `q:"ecStatus"`
	NetworkId    string         `q:"networkId"`
	Visible      bool           `q:"visible"`
	Region       string         `q:"region"`
}

type GetOpts struct {
	ServerID string
}

func (o *GetOpts) Verify() error {
	if o.ServerID == "" {
		return errors.ErrMissingInput{Argument: "GetOpts.ServerID"}
	}
	return nil
}

type UpdateOpts struct {
	ServerID string `json:"id"`
	// update server password
	Password string `json:"password,omitempty"`
	// update server name
	Name string `json:"name,omitempty"`
}

func (opts *UpdateOpts) Verify() error {
	if opts.ServerID == "" {
		return errors.ErrMissingInput{Argument: "UpdateOpts.ServerID"}
	}
	var count = 0
	if opts.Name != "" {
		count++
	}
	if opts.Password != "" {
		count++
	}
	if count != 1 {
		err := errors.ErrInvalidInput{}
		err.Info = fmt.Sprintf("vaild UpdateOpts,only support to update name or password once,UpdateOpts is:%v", opts)
		return err
	}
	return nil
}

type DeleteOpts struct {
	ServerID string
}

func (o *DeleteOpts) Verify() error {
	if o.ServerID == "" {
		return errors.ErrMissingInput{Argument: "DeleteOpts.ServerID"}
	}
	return nil
}

type ActionOpts struct {
	ServerID string
	Action   ActionType
}

func (o *ActionOpts) Verify() error {
	if o.ServerID == "" {
		return errors.ErrMissingInput{Argument: "ActionOpts.ServerID"}
	}
	if o.Action == "" {
		return errors.ErrMissingInput{Argument: "ActionOpts.Action"}
	}
	return nil
}

type RebuildOpts struct {
	ServerID  string `json:"serverId" required:"true"`
	ImageID   string `json:"imageId,omitempty"`
	AdminPass string `json:"adminPass,omitempty"`
	UserData  string `json:"userData,omitempty"`
}

func (o *RebuildOpts) Verify() error {
	if o.ServerID == "" {
		return errors.ErrMissingInput{Argument: "RebuildOpts.ServerID"}
	}
	if o.ImageID == "" {
		return errors.ErrMissingInput{Argument: "RebuildOpts.ImageID"}
	}
	return nil
}
