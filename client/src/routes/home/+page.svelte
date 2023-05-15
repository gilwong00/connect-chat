<script lang="ts">
  import { onMount } from 'svelte';
  import { roomClient } from '../../clients';
  import { getCookie } from 'svelte-cookie';
  import { roomStore } from '../../store';
  import { goto } from '$app/navigation';
  import { promiseHandler } from '../../utils';
  import type { GetAllRoomsResponse } from '../../gen/proto/room_pb';
  import type { ConnectError } from '@bufbuild/connect';
  import AvailableRoom from '../../components/AvailableRoom.svelte';

  let fetchRoomsError: boolean = false;

  onMount(async () => {
    // fetch room if auth
    const authToken = getCookie('token');
    if (!authToken.length) return goto('/login');
    const [res, err] = await promiseHandler<
      GetAllRoomsResponse,
      ConnectError | null
    >(
      roomClient.getAllRooms(
        {},
        {
          headers: {
            Authorization: `Bearers ${authToken}`
          }
        }
      )
    );
    if (err !== null) {
      fetchRoomsError = true;
    }
    if (res !== null) {
      roomStore.setRooms(res.rooms);
    }
  });
</script>

<div class="home-container">
  <div class="home-header">
    <h2>Connect Rooms</h2>
    <button class="create-room-btn" on:click={() => goto('/room/create')}
      >Create room</button
    >
  </div>
  {#if fetchRoomsError}
    <h2>Failed to fetch rooms</h2>
  {/if}
  <div class="room-list">
    {#if $roomStore.rooms.length}
      {#each $roomStore.rooms as room (room.roomId)}
        <AvailableRoom roomId={room.roomId} roomName={room.roomName} />
      {/each}
    {/if}
  </div>
</div>

<style lang="scss">
  .home-container {
    display: flex;
    flex-direction: column;
    margin: 0 auto;
    background: #fff;
    width: 100%;
    height: 100%;
    padding-bottom: 30px;
    padding-left: 10px;
    padding-right: 10px;
    border-radius: 4px;
  }
  h2 {
    text-align: center;
    font-size: 30px;
  }

  .home-header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding-left: 10px;
    padding-right: 10px;
  }
  .create-room-btn {
    padding: 12px 10px;
    background: purple;
    color: #fff;
    border-radius: 4px;
    border: none;
    cursor: pointer;
  }

  .room-list {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    flex-wrap: wrap;
    padding: 10px;
  }
</style>
