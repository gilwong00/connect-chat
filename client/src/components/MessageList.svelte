<script lang="ts">
  import type { Message } from '../@types';

  export let messages: Array<Message>;
</script>

<div>
  {#each messages as message}
    <div class={message.type === 'self' ? 'right_align' : 'left-align'}>
      <div class="message_column">
        <div class="username">{message.username}</div>
        <div
          class={`message_content ${
            message.type === 'self' ? 'self_content' : 'receive_content'
          }`}
        >
          {message.content}
        </div>
      </div>
    </div>
  {/each}
</div>

<style lang="scss">
  .message_column {
    flex-direction: column;
    padding: 5px 0;
  }
  .username {
    color: #000;
    font-size: 12px;
    margin-bottom: 1px;
  }
  .message_content {
    display: inline-block;
    padding: 5px 10px;
    width: auto;
    border-radius: 20px;
    font-size: 12px;
  }

  .self_content {
    background-color: blue;
    color: #fff;
  }

  .receive_content {
    background-color: darkgray;
    color: #fff;
  }

  .right_align {
    text-align: right;
  }

  .left_align {
    /* background-color: lightgray; */
    text-align: left;
  }
</style>
