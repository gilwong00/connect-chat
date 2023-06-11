<script lang="ts">
  import { page } from '$app/stores';
  import { onMount } from 'svelte';
  import { getCookie } from 'svelte-cookie';
  import { roomClient } from '../../../clients';
  import { Card, Heading } from 'flowbite-svelte';
  import ChatInput from '../../../components/ChatInput.svelte';
  import { channelStore } from '../../../store/channelStore';
  import { goto } from '$app/navigation';

  let roomMembers;
  $: messages = [];

  const channel = $channelStore.channels[$page.params.roomId];

  onMount(async () => {
    if (!channel) {
      return goto('/');
    }

    try {
      const token = getCookie('token');
      // fetch people in room
      const res = await roomClient.getRoomMembers(
        { roomId: $page.params.roomId },
        {
          headers: {
            Authorization: `Bearer ${token}`
          }
        }
      );
      roomMembers = res.members;
    } catch (err) {
      // do something with error
      console.log(err);
      return goto('/');
    }
  });

  channel.onmessage = (incomingMessage: unknown) => {};

  const sendMessage = async (message: string) => {
    channel.send(message);
  };
</script>

<div class="room-container">
  <Heading class="mb-4" tag="h3">Room: {$page.params.roomId}</Heading>
  <Card size="xl">
    <div class="chat-container">
      <div class="messages-container" />
      <ChatInput />
    </div>
  </Card>
</div>

<style lang="scss">
  .room-container {
    margin: 0 auto;
  }

  .messages-container {
    overflow-y: auto;
    height: 700px;
  }

  .chat-container {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    width: 800px;
    height: 500px;
  }
</style>
