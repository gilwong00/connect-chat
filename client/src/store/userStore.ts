import { writable } from 'svelte/store';
import type { User } from '../gen/proto/user_pb';

type State = {
  user: User | null;
};

export const userStore = (() => {
  const { subscribe, set, update } = writable<State>({
    user: null
  });

  const setUser = (user: User) => {
    update(state => ({
      ...state,
      user
    }));
  };

  return {
    subscribe,
    set,
    update,
    setUser
  };
})();
