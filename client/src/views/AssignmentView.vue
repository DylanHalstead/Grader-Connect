<template>
  <div class="mx-4 mb-8">
    <h1 class="mb-6 text-white font-bold sm:text-4xl sm:text-left text-3xl text-center">
      Awaiting Grading Assignments
    </h1>
    <div class="grid gap-4 sm:grid-cols-1 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5">
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
    </div>
  </div>
  <div class="mx-4 mb-8">
    <h1 class="mb-6 text-white font-bold sm:text-4xl sm:text-left text-3xl text-center">
      Graded Assignments
    </h1>
    <div class="grid gap-4 sm:grid-cols-1 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5">
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
      <AssignmentCard />
    </div>
  </div>
</template>

<script setup>
import AssignmentCard from '../components/AssignmentCard.vue'
</script>
