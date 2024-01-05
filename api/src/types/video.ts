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
  id: string;
  supplierId: string;
  web_url: string;
  title: string;
  description: string;
  video_id: string;
  thumbnail_url: string;
}

export interface VideoDTO {
  video: Video;
  youtube: {
    id: string;
    title: string;
    thumbnail: string;
    description: string;
  };
}
