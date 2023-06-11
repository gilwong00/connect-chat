<script lang="ts">
  import { goto } from '$app/navigation';
  import { WEBSOCKET_URL } from '../clients';
  import { userStore } from '../store';
  import { Card, Button, Input } from 'flowbite-svelte';
  import { channelStore } from '../store/channelStore';

  export let roomName: string = '';
  export let roomId: string = '';

  const user = $userStore.user;

  const joinRoom = async () => {
    if (user) {
      const ws = new WebSocket(
        `${WEBSOCKET_URL}/room/join?roomId=${roomId}&clientId=${user.id}&username=${user.username}`
      );
      if (ws.OPEN) {
        // save connection with room Id in store
        channelStore.setChannel(roomId, ws);
        return goto(`/room/${roomId}`);
      }
    }
  };
</script>

<div class="room-item-container">
  <Card>
    <div class="room-details-container">
      <div class="room-details">
        <div>Room</div>
        <p>{roomName}</p>
      </div>
      <Button on:click={joinRoom}>Join</Button>
    </div>
  </Card>
</div>

<style lang="scss">
  .room-item-container {
    width: 300px;
    margin-top: 15px;
    margin-bottom: 15px;
  }

  .room-details-container {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;

    > button {
      border-radius: 4px;
      border: none;
      background-color: purple;
      color: #fff;
      padding: 5px;
      width: 60px;
      cursor: pointer;
    }
  }

  .room-details {
    display: flex;
    flex-direction: column;

    > p {
      color: blue;
      margin-top: 10px;
    }
  }
</style>
