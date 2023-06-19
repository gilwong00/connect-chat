<script lang="ts">
  import { page } from '$app/stores';
  import { onMount } from 'svelte';
  import { getCookie } from 'svelte-cookie';
  import { roomClient } from '../../../clients';
  import { Card, Heading } from 'flowbite-svelte';
  import ChatInput from '../../../components/ChatInput.svelte';
  import { channelStore } from '../../../store/channelStore';
  import { goto } from '$app/navigation';
  import { userStore } from '../../../store';
  import type { Message } from '../../../@types';
  import MessageList from '../../../components/MessageList.svelte';

  const currentUser = $userStore.user;

  let messages = [] as Array<Message>;
  let roomMembers = [] as Array<{ username: string }>;

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
      roomMembers = res.members.map(member => ({
        usernanme: member.memberName
      })) as unknown as Array<{ username: string }>;
    } catch (err) {
      // do something with error
      console.log(err);
      return goto('/');
    }
  });

  if (channel) {
    channel.onmessage = (incomingMessage: MessageEvent) => {
      const message: Message = JSON.parse(incomingMessage.data);
      if (message.content === 'A new user has joined the room') {
        // append user to chat room
        roomMembers = [...roomMembers, { username: message.username }];
      }

      if (message.content === 'user left the chat') {
        // remove user
        const deleteUser = roomMembers.filter(
          user => user.username !== message.username
        );
        roomMembers = [...deleteUser];
        messages = [...messages, message];
        return;
      }

      if (currentUser?.username === message.username) {
        message.type = 'self';
      } else {
        message.type = 'receive';
      }
      messages = [...messages, message];
    };
    // at some point on close should remove the user from the room
    channel.close = () => {};
  }

  const sendMessage = async (message: string) => {
    if (channel === null) {
      return goto('/');
    }

    channel.send(message);
  };
</script>

<div class="room-container">
  <Heading class="mb-4" tag="h3">Room: {$page.params.roomId}</Heading>
  <Card size="xl">
    <div class="chat-container">
      <div class="messages-container">
        <MessageList {messages} />
      </div>
      <ChatInput handleSendMessage={sendMessage} />
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
