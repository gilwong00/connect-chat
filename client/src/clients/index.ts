import { createPromiseClient } from '@bufbuild/connect';
import { createConnectTransport } from '@bufbuild/connect-web';
import { UserService } from '../gen/proto/user_connect';
import { RoomService } from '../gen/proto/room_connect';

const BASE_URL = 'http://localhost:8080';

export const userClient = createPromiseClient(
  UserService,
  createConnectTransport({
    baseUrl: BASE_URL
  })
);

export const roomClient = createPromiseClient(
  RoomService,
  createConnectTransport({
    baseUrl: BASE_URL
  })
);
