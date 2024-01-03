import { FastifyInstance } from 'fastify';
import VideoService from 'services/videos';
import VideoController from 'controllers/videos';

const videos = async (fastify: FastifyInstance) => {
  const videoService = new VideoService(fastify);
  const videoController = new VideoController(videoService);

  fastify.get('/videos/:id', videoController.getVideos.bind(videoController));
  fastify.post('/videos', videoController.createVideo.bind(videoController));

  fastify.put('/videos/:id', videoController.updateVideo.bind(videoController));
  fastify.delete('/videos/:id', videoController.deleteVideo.bind(videoController));
};

export default videos;
