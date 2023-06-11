<script lang="ts">
  import { browser } from '$app/environment';
  import { goto } from '$app/navigation';
  import { getCookie } from 'svelte-cookie';
  import { promiseHandler } from '../../../utils';
  import type { ConnectError } from '@bufbuild/connect';
  import { roomClient } from '../../../clients';
  import { roomStore } from '../../../store';
  import type { CreateRoomResponse, Room } from '../../../gen/proto/room_pb';
  import { Card, Button, Input } from 'flowbite-svelte';

  let roomName: string = '';

  const authToken = browser ? getCookie('token') : null;

  const returnHome = () => goto('/');

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
  <Card>
    <h3 class="text-center">Create new room</h3>
    <div class="my-6">
      <Input
        type="text"
        name="new-room"
        bind:value={roomName}
        placeholder="Room name"
        on:change={handleNameChange}
        required
      />
    </div>
    <Button class="mb-4" color="purple" on:click={handleCreateRoom}
      >Create</Button
    >
    <Button color="red" on:click={returnHome}>Cancel</Button>
  </Card>
</div>

<style lang="scss">
  .create-room-container {
    align-items: center;
    margin: auto;
    width: 500px;
  }
</style>
