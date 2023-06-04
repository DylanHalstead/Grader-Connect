<template>
  <div class="flex min-h-full flex-1 flex-col justify-center mx-4">
    <h2 class="text-center text-2xl font-bold tracking-tight text-white">
      Sign in to your account
    </h2>
    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <form class="space-y-6" action="#" method="POST">
        <div>
          <label for="email" class="block text-sm font-medium leading-6 text-white"
            >Email address</label
          >
          <div class="mt-2">
            <input
              id="email"
              name="email"
              type="email"
              autocomplete="email"
              required=""
              class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm focus:ring-transparent focus:shadow-[0_0_10px_1px_rgba(22,163,74,1)] sm:text-sm sm:leading-6 transition-shadow"
            />
          </div>
        </div>
        <div>
          <div class="flex items-center justify-between">
            <label for="password" class="block text-sm font-medium leading-6 text-white"
              >Password</label
            >
          </div>
          <div class="mt-2">
            <input
              id="password"
              name="password"
              type="password"
              autocomplete="current-password"
              required=""
              class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm focus:ring-transparent focus:shadow-[0_0_10px_1px_rgba(22,163,74,1)] sm:text-sm sm:leading-6 transition-shadow"
            />
          </div>
        </div>
        <div>
          <button
            type="submit"
            class="flex w-full justify-center rounded-md bg-emerald-600 hover:bg-emerald-500 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-emerald-600 hover:shadow-[0_0_12px_0_rgba(22,163,74,0.5)] duration-300 transition-all ease-in-out"
          >
            Sign in
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup></script>
