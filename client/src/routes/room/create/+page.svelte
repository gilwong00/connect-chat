<script lang="ts">
  import { browser } from '$app/environment';
  import { goto } from '$app/navigation';
  import { getCookie } from 'svelte-cookie';
  import { promiseHandler } from '../../../utils';
  import type { ConnectError } from '@bufbuild/connect';
  import { roomClient } from '../../../clients';
  import { roomStore } from '../../../store';
  import type { CreateRoomResponse, Room } from '../../../gen/proto/room_pb';

  let roomName: string = '';

  const authToken = browser ? getCookie('token') : null;

  const returnHome = () => goto('/home');

  const handleNameChange = (e: Event) => {
    roomName = (e.target as HTMLInputElement).value;
  };

  const handleCreateRoom = async () => {
    const [res, err] = await promiseHandler<CreateRoomResponse, ConnectError>(
      roomClient.createRoom(
        { name: roomName },
        {
          headers: {
            Authorization: `Bearers ${authToken}`
          }
        }
      )
    );

    if (err !== null) {
      // show error
    }

    if (res !== null) {
      roomStore.addNewRoom({
        roomId: res.roomId,
        roomName: res.name
      } satisfies Pick<Room, 'roomId' | 'roomName'>);
      returnHome();
    }
  };
</script>

<div class="create-room-container">
  <h3>Create new room</h3>
  <input
    class="create-room-input"
    placeholder="Room Name"
    bind:value={roomName}
    on:change={handleNameChange}
  />
  <div class="action-row">
    <button class="cancel" on:click={returnHome}>Cancel</button>
    <button on:click={handleCreateRoom}>Create</button>
  </div>
</div>

<style lang="scss">
  .create-room-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    background: #fff;
    width: 500px;
    height: 150px;
    margin: auto;
    padding: 20px;
    border-radius: 8px;
    margin-top: 100px;
    padding-bottom: 50px;
  }

  .create-room-input {
    padding: 10px;
    width: 350px;
    border: 1px solid lightgrey;
    border-radius: 4px;
  }

  .action-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 380px;
    margin: auto;
    margin-top: 20px;

    > button {
      padding: 10px 10px;
      background: purple;
      color: #fff;
      border-radius: 4px;
      border: none;
      cursor: pointer;
      width: 120px;
    }

    .cancel {
      background: red;
    }
  }
</style>
