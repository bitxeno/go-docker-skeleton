<template>
  <div class="app-wrapper">
    <div class="navbar bg-base-100">
      <div class="flex-1">
        <a href="https://daisyui.com/" class="btn btn-ghost normal-case text-xl"
          >daisyUI</a
        >
      </div>
      <div class="flex-none">
        <ul class="menu menu-horizontal px-1">
          <li><router-link to="/about">About</router-link></li>
        </ul>
      </div>
    </div>

    <div class="main-container">
      <router-view />
    </div>
  </div>
</template>
  
  <script>
export default {
  name: "App",
};
</script>