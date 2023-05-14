<script lang="ts">
  import { onMount } from 'svelte';
  import { roomClient } from '../../clients';
  import { getCookie } from 'svelte-cookie';
  import { roomStore } from '../../store';
  onMount(async () => {
    // fetch room if auth
    const authToken = getCookie('token');
    if (authToken || true) {
      const res = await roomClient.getAllRooms({});
      roomStore.setRooms(res.rooms);
    }
  });
</script>

<div>
  <h2>Rooms</h2>
  {#if $roomStore.rooms.length}
    {#each $roomStore.rooms as room (room.roomId)}
      <div>{room.roomName}</div>
    {/each}
  {/if}
</div>

<style lang="scss"></style>
