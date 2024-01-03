import { FastifyReply, FastifyRequest } from 'fastify';
import TimelineService from 'services/timelines';
import { Timeline } from 'types/about';

class TimelineController {
  private timelineService: TimelineService;

  constructor(timelineService: TimelineService) {
    this.timelineService = timelineService;
  }

  async getTimelines(request: FastifyRequest, reply: FastifyReply) {}

  async createTimeline(request: FastifyRequest<{ Body: Timeline }>, reply: FastifyReply) {}

  async updateTimeline(request: FastifyRequest<{ Params: { id: string }; Body: Timeline }>, reply: FastifyReply) {}

  async deleteTimeline(request: FastifyRequest<{ Params: { id: string } }>, reply: FastifyReply) {}
}

export default TimelineController;
