import { writable } from 'svelte/store';

type State = {
  channels: Record<string, WebSocket>;
};

export const channelStore = (() => {
  const { subscribe, set, update } = writable<State>({
    channels: {}
  });

  const setChannel = (roomId: string, channel: WebSocket) => {
    update(state => ({
      ...state,
      channels: {
        ...state.channels,
        [roomId]: channel
      }
    }));
  };

  return {
    subscribe,
    set,
    setChannel
  };
})();
