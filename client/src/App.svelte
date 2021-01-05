<script lang="ts">
  import { formatDistanceToNow } from "date-fns";
  import { onMount } from "svelte";
  import { slide } from "svelte/transition";
  import Block from "./Block.svelte";
  import Button from "./Button.svelte";
  import { blocks } from "./store";

  let isCollapsed = false;
  let lastSaved = new Date();

  function addBlock() {
    $blocks = [
      ...$blocks,
      {
        header: new Date().toLocaleDateString(),
        body: "",
        id: Math.random()
          .toString(36)
          .substr(2, 9)
      }
    ];
    updateBlocks();
  }

  async function fetchBlocks() {
    const response = await fetch("/api/blocks");
    const data = await response.json();
    blocks.update(() => data);
  }

  async function updateBlocks() {
    const response = await fetch("/api/blocks/update", {
      mode: "no-cors",
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(blocks)
    });
  }

  function toggleAll() {
    isCollapsed = !isCollapsed;
  }

  const debounce = (func, wait) => {
    let timeout;

    return function executedFunction(...args) {
      const later = () => {
        timeout = null;
        func(...args);
      };
      clearTimeout(timeout);
      timeout = setTimeout(later, wait);
    };
  };

  onMount(() => {
    fetchBlocks();
  });
</script>



<main class="container w-1/3">
  <div class="prose py-3 max-w-none">
    <div class="flex justify-between items-center">
      <div>
        <h1>Simple note</h1>
        <p class="text-opacity-75">
          {#await updateBlocks}
            Saving...
          {:then number}
            Saved {formatDistanceToNow(lastSaved)}
          {:catch error}
            <p style="color: red">{error}</p>
          {/await}
        </p>
      </div>

      <div class="space-x-3">
        <Button onClick={toggleAll}>CA</Button>
        <Button onClick={addBlock}>A</Button>
      </div>
    </div>
    {#each blocks as block (block.id)}
      <div transition:slide>
        <Block
          onKeyUp={debounce(updateBlocks, 400)}
          onBlur={updateBlocks}
          {isCollapsed}
          header={block.header}
          body={block.body} />
      </div>
    {:else}
      <div>Nothing here!</div>
    {/each}
  </div>
</main>






<style global lang="postcss">

  /* only apply purgecss on utilities, per Tailwind docs */
  /* purgecss start ignore */
  @tailwind base;
  @tailwind components;
  /* purgecss end ignore */

  @tailwind utilities;

</style>
