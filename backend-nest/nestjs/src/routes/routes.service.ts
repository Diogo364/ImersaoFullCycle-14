import { Injectable } from '@nestjs/common';
import { CreateRouteDto } from './dto/create-route.dto';
import { PrismaService } from 'src/prisma/prisma/prisma.service';

@Injectable()
export class RoutesService {
  constructor(
    private prismaService: PrismaService 
  ) { }

  create(createRouteDto: CreateRouteDto) {
    return this.prismaService.route.create({
      data: {
        name: createRouteDto.name,
        source: {
          name: createRouteDto.source.name,
          location: {
            lat: createRouteDto.source.location.lat,
            lng: createRouteDto.source.location.lng,
          },
        },
        destination: {
          name: createRouteDto.destination.name,
          location: {
            lat: createRouteDto.destination.location.lat,
            lng: createRouteDto.destination.location.lng,
          },
        },
      }
    })
  }

  findAll() {
    return this.prismaService.route.findMany();
  }
}
