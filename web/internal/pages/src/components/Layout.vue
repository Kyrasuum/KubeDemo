<script setup lang="ts">
</script>

<template>
	<q-layout view="hHh LpR fFf">
		<q-header elevated class="bg-grey-9 text-grey" style="opacity: 95%">
			<section class="row items-center">
				<q-space />

				<span class="px-3">|</span>
				<router-link to="/" active-class="text-white">Home</router-link>
				<span class="px-3">|</span>
				<router-link to="about" active-class="text-white">About</router-link>
				<span class="px-3">|</span>
				<q-btn flat no-caps color="grey-5" label="Dark Mode" @click="toggleDarkMode()" />

				<q-space />
			</section>
		</q-header>

		<q-page-container>
			<q-page class="row" :class="[$q.dark.isActive ? 'dark' : '']">
				<div class="mx-auto w-5/6 column">
				    <slot name="default"/>
				</div>
			</q-page>
		</q-page-container>

		<q-footer elevated class="bg-grey-9 text-grey">
			<q-toolbar> </q-toolbar>
		</q-footer>
	</q-layout>
</template>

<style scoped lang="scss">
</style>
