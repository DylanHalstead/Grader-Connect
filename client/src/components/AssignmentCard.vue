<template>
  <div class="relative">
    <!-- TODO: Add function where if modal is active add hover effects -->
    <div
      class="p-4 bg-slate-950 text-white border rounded-lg border-transparent cursor-pointer hover:border-emerald-600 hover:border-opacity-40 transition-all ease-in-out duration-500 hover:shadow-[0_4px_8px_rgba(22,163,74,.2)] transform hover:-translate-y-2"
      @click="toggleModal"
    >
      <div class="w-full pb-4 flex flex-col items-center md:items-start">
        <h3 class="font-semibold text-2xl">Title Goes Here</h3>
        <p class="italic font-medium text-sm">Due: 12/12/2023</p>
      </div>
      <!-- TODO: fill as people complete grading -->
      <p
        class="w-full py-1 text-md text-center bg-emerald-600 bg-opacity-20 border border-emerald-600 rounded-md border-opacity-70"
      >
        x/y done grading
      </p>
    </div>
    <AssignmentModal :modalActive="modalActive" @close-modal="toggleModal" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import AssignmentModal from './AssignmentModal.vue'

const modalActive = ref(null)
const toggleModal = () => {
  modalActive.value = !modalActive.value
}
</script>
