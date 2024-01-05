package models

type YoutubeVideo struct {
    ID      string `json:"id"`
    Snippet struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        Thumbnails  struct {
            Medium struct {
                URL string `json:"url"`
            } `json:"medium"`
        } `json:"thumbnails"`
    } `json:"snippet"`
}

type YoutubeApiResponse struct {
    Data struct {
        Items []YoutubeVideo `json:"items"`
    } `json:"data"`
}

type Video struct {
    ID            string `json:"id"`
    SupplierID    string `json:"supplierId"`
    WebURL        string `json:"web_url"`
    Title         string `json:"title"`
    Description   string `json:"description"`
    VideoID       string `json:"video_id"`
    ThumbnailURL  string `json:"thumbnail_url"`
}