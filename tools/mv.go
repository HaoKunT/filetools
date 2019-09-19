package tools

type MoveOptions struct {
	Src string
	Dst string
}

var moveOptions *MoveOptions

func Move(options *MoveOptions) error {
	moveOptions = options
	extInfo, err := NewExtInfo(options.Src)
	if err != nil {
		return err
	}

}
