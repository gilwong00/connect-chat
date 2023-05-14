import { writable } from 'svelte/store';

type State = {
  rooms: Array<{ roomId: string; roomName: string }>;
};

export const store = writable<State>({
  rooms: []
});
