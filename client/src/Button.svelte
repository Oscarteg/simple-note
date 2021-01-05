<script lang="ts">
  export let onClick;
</script>

<button
  on:click|preventDefault={onClick}
  class="hover:bg-gray-800 hover:text-white text-gray-800 font-bold py-1 px-2
  focus:outline-none rounded transition duration-300 ease-in-out">
  <slot />
</button>
