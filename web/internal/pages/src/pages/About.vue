<script setup lang="ts">
import layout from 'src/components/Layout.vue';
</script>

<template>
	<layout>
    	<template #default>
        	<div class="col w-full column justify-center">
        		<div class="flex justify-center">
        			<q-card class="q-pa-md">
        				<q-card-section>
        					<div class="absolute-bottom">
        						<div class="text-h6">KubeDemo</div>
        						<div class="text-subtitle2">by: Phillip Lorentz</div>
        					</div>
        				</q-card-section>

        				<q-card-section>
        					Kubernetes demo application
        				</q-card-section>

        				<q-card-actions>
        					<q-space/>
        					<q-btn flat>Contact Us</q-btn>
        					<q-space/>
        				</q-card-actions>
        			</q-card>
        		</div>
        	</div>
    	</template>
	</layout>
</template>

<style scoped>
</style>
