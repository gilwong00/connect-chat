export type MessageType = 'self' | 'receive';

export type Message = {
  content: string;
  clientId: string;
  username: string;
  roomId: string;
  type: MessageType;
};
