export class CreateRouteDto {
  name: string
  source: PlaceDto
  destination: PlaceDto
}

export class PlaceDto {
  name: string
  location: CoordDto
}

export class CoordDto {
  lat: number
  lng: number
}