<script lang="ts">
  import { userStore } from '../store';

  export let roomName: string = '';
  export let roomId: string = '';

  const user = $userStore.user;
  console.log(user);

  const joinRoom = async () => {
    console.log('join click');
  };
</script>

<div class="room-item-container">
  <div class="room-details-container">
    <div class="room-details">
      <label>Room</label>
      <p>{roomName}</p>
    </div>
    <button on:click={joinRoom}>Join</button>
  </div>
</div>

<style lang="scss">
  .room-item-container {
    height: 65px;
    width: 250px;
    border: 1px solid lightgray;
    padding: 15px;
    border-radius: 10px;
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
