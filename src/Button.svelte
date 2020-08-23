<script lang="ts">
  export let onClick;
</script>

<button
  on:click|preventDefault={onClick}
  class="hover:bg-white hover:text-gray-800 text-white font-bold py-1 px-2
  focus:outline-none rounded">
  <slot />
</button>
