<script lang="ts">
  import { Alert, Textarea, ToolbarButton } from 'flowbite-svelte';

  export let handleSendMessage = (message: string) => {};

  let message: string;

  const handleMessageChange = (e: Event) => {
    message = (e.target as HTMLInputElement).value;
  };

  const sendMessage = () => {
    handleSendMessage(message);
    message = '';
  };
</script>

<div>
  <label for="chat" class="sr-only">Your message</label>
  <Alert color="dark" class="px-3 py-2">
    <svelte:fragment slot="icon">
      <Textarea
        id="chat"
        class="mx-4"
        rows="1"
        placeholder="Say something nice"
        value={message}
        on:change={handleMessageChange}
      />
      <ToolbarButton
        type="submit"
        color="blue"
        class="rounded-full text-primary-600 dark:text-primary-500"
        on:click={sendMessage}
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="w-6 h-6"
          ><path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M6 12L3.269 3.126A59.768 59.768 0 0121.485 12 59.77 59.77 0 013.27 20.876L5.999 12zm0 0h7.5"
          />
          <span class="sr-only">Send message</span>
        </svg>
      </ToolbarButton>
    </svelte:fragment>
  </Alert>
</div>
