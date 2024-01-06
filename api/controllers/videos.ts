import VideoService from 'services/videos';

import { FastifyReply, FastifyRequest } from 'fastify';

import { youtube as YouTube } from '@googleapis/youtube';
import { YoutubeApiResponse, Video } from 'types/video';
import secrets from 'utils/secrets';

class VideoController {
  private videoService: VideoService;

  constructor(videoService: VideoService) {
    this.videoService = videoService;
  }

  async getVideos(request: FastifyRequest<{ Params: { id: string } }>, reply: FastifyReply) {
    const id = request.params.id;

    try {
      const videos = await this.videoService.getVideos(id);
      return reply.code(200).send(videos);
    } catch (error) {
      console.error('Controller Error:', error);

      return reply.code(500).send({ error: 'Internal Server Error' });
    }
  }

  async createVideo(request: FastifyRequest<{ Body: Video }>, reply: FastifyReply) {
    const data = request.body;

    try {
      const videoData = await this.transformData(data);

      const videoResponse = await this.videoService.createVideo(videoData);

      return reply.code(201).send(videoResponse);
    } catch (error) {
      console.error('Controller Error:', error);

      return reply.code(500).send({ error: 'Internal Server Error' });
    }
  }

  async updateVideo(request: FastifyRequest<{ Params: { id: string }; Body: Video }>, reply: FastifyReply) {
    const data = request.body;
    const id = request.params.id;

    try {
      const videoData = await this.transformData(data);

      const videoResponse = await this.videoService.updateVideo(id, videoData);

      return reply.code(200).send(videoResponse);
    } catch (error) {
      console.error('Controller Error:', error);

      return reply.code(500).send({ error: 'Internal Server Error' });
    }
  }

  async deleteVideo(request: FastifyRequest<{ Params: { id: string } }>, reply: FastifyReply) {
    const id = request.params.id;

    try {
      const video = this.videoService.deleteVideo(id);

      return reply.code(200).send({
        video,
        message: 'Supplier deleted',
      });
    } catch (error) {
      console.error('Controller Error:', error);

      return reply.code(500).send({ error: 'Internal Server Error' });
    }
  }

  async transformData(data: Video) {
    const { web_url } = data;
    const videoId = web_url.split('v=')[1].split('&')[0];
    try {
    } catch (error) {
      console.error('Controller Error:', error);
    }

    const youtube = YouTube({
      version: 'v3',
      auth: secrets.youtube_api_key,
    });

    const videoResponse = (await youtube.videos.list({
      part: ['id', 'snippet'],
      id: videoId,
      maxResults: 1,
    })) as YoutubeApiResponse;

    if (!videoResponse.data.items || videoResponse.data.items.length === 0) {
      throw new Error('Video not found on YouTube');
    }

    const id = videoResponse.data.items[0].id;
    const thumbnail = videoResponse.data.items[0].snippet.thumbnails.medium.url;

    const { title, description } = videoResponse.data.items[0].snippet;

    const videoData = {
      video: {
        ...data,
      },
      youtube: {
        id,
        title,
        thumbnail,
        description,
      },
    };

    return videoData;
  }
}

export default VideoController;
