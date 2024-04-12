<script lang="ts">
    import * as AlertDialog from '$lib/components/ui/alert-dialog/index.js';

    export let show = false;
</script>

<AlertDialog.Root bind:open={show}>
    <AlertDialog.Content>
        <AlertDialog.Header>
            <AlertDialog.Title class="text-4xl font-bold">
                You discovered a <span class="text-orange-400 font-black">Premium</span> feature!
            </AlertDialog.Title>
            <AlertDialog.Description class="text-xl">
                With StoryQuiz Premium you can enjoy even more quiz exporting fun!
            </AlertDialog.Description>
        </AlertDialog.Header>
        <AlertDialog.Footer>
            <AlertDialog.Action class="font-bold">Awesome</AlertDialog.Action>
        </AlertDialog.Footer>
    </AlertDialog.Content>
</AlertDialog.Root>
