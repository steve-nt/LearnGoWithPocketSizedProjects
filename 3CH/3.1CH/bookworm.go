f, err := os.Open(filePath)
if err != nil {
    return nil, err
}
defer f.Close()