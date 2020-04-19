package pud

// 判断文件大小是否合适
func (p *PhotoUploader) PhotoSizeIsValid(photo []byte, size int) (bool, error) {
	return len(photo) <= size, nil
}
