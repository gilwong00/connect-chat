import { writable } from 'svelte/store';
import type { Room } from '../gen/proto/room_pb';

type State = {
  rooms: Array<{ roomId: string; roomName: string }>;
};

export const roomStore = (() => {
  const { subscribe, set, update } = writable<State>({
    rooms: []
  });

  const setRooms = (rooms: Array<Room>) => {
    update(state => ({
      ...state,
      rooms
    }));
  };

  const addNewRoom = (room: Room) => {
    update(state => ({
      ...state,
      rooms: [...state.rooms, room]
    }));
  };

  return {
    subscribe,
    set,
    update,
    setRooms,
    addNewRoom
  };
})();
