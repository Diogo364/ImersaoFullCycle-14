import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { RoutesModule } from './routes/routes.module';
import { RoutesController } from './routes/routes.controller';
import { RoutesService } from './routes/routes.service';
import { PrismaModule } from './prisma/prisma.module';

@Module({
  imports: [
    RoutesModule,
    PrismaModule
  ],
  controllers: [
    AppController,
    RoutesController
  ],
  providers: [
    AppService,
    RoutesService
  ],
})
export class AppModule {}
