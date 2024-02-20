export interface YoutubeVideo {
    id: string;
    snippet: {
        title: string;
        description: string;
        thumbnails: {
            medium: {
                url: string;
            };
        };
    };
}

export interface YoutubeApiResponse {
    data: {
        items: YoutubeVideo[];
    };
}

export interface Video {
    id?: string;
    supplier_id: string;
    web_url: string;
    title: string;
    description: string;
    video_id: string;
    thumbnail_url: string;
    created?: string;
}

export interface VideoWebUrl {
    supplier_id: string;
    web_url: string;
}
