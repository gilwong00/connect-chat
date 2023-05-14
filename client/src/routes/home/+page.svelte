<script lang="ts">
  import { onMount } from 'svelte';
  import { roomClient } from '../../clients';
  import { getCookie } from 'svelte-cookie';
  import { roomStore } from '../../store';

  let fetchRoomsError: boolean = false;

  onMount(async () => {
    try {
      // fetch room if auth
      const authToken = getCookie('token');
      if (authToken || true) {
        const res = await roomClient.getAllRooms(
          {},
          {
            headers: {
              Authorization: `Bearers ${authToken}`
            }
          }
        );
        roomStore.setRooms(res.rooms);
      }
    } catch (err) {
      fetchRoomsError = true;
    }
  });
</script>

<div>
  <h2>Rooms</h2>
  {#if fetchRoomsError}
    <h2>Failed to fetch rooms</h2>
  {/if}
  {#if $roomStore.rooms.length}
    {#each $roomStore.rooms as room (room.roomId)}
      <div>{room.roomName}</div>
    {/each}
  {/if}
</div>

<style lang="scss"></style>
