package downloader

type Repository interface {
	DownloadCSV(url string) ([][]string, error)
}
