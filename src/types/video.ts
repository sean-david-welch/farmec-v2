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
