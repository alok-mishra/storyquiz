<script lang="ts">
    import * as AlertDialog from '$lib/components/ui/alert-dialog/index.js';
    import { Button } from '$lib/components/ui/button/index.js';
    export let fileTypes: string[] = [],
        onDrop: (file: File, resultText: string) => void,
        dropColor: string = 'bg-stone-400';

    export let disabled: boolean = false;

    let showPremiumDialog = false;

    const fileTypeMap = {
        // file type map of allowed file types
        csv: 'text/csv',
        docx: 'application/vnd.openxmlformats-officedocument.wordprocessingml.document',
        xls: 'application/vnd.ms-excel',
        xlsm: 'application/vnd.ms-excel.sheet.macroEnabled.12',
        xlsx: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
    };

    function handleDrop(event: DragEvent): void {
        event.preventDefault();
        (event.target as HTMLElement).classList.remove(dropColor);

        if (disabled) {
            showPremiumDialog = true;
            return;
        }

        const file = event.dataTransfer.files[0];

        console.log('file', file);

        console.log('disabled:', disabled);

        const fileType = file.name.split('.').pop();

        console.log('filetype:', fileType);

        // const allowedTypes = fileTypes.map((type) => fileTypeMap[type] || type);

        // console.log('allowed:', allowedTypes);

        // if (allowedTypes.includes(file.type) || allowedTypes.includes(fileType)) {
        if (fileTypes.includes(fileType as string)) {
            onDrop(file, `Exporting: ${file.name}`);
        } else {
            onDrop(null, `<span class='text-4xl'>😕</span> Seriously?! Please only use ${fileTypes.join(', ')} files!`);
        }
    }
</script>

<article
    id="drop"
    class="border-2 border-dotted p-6 m-4 rounded-lg bg-opacity-60 max-w-60"
    on:drop={handleDrop}
    on:dragover={(e) => e.preventDefault()}
    on:dragenter={(e) => {
        e.preventDefault();
        e.target.classList.add(dropColor);
    }}
    on:dragleave={(e) => {
        e.preventDefault();
        // e.target.classList.remove('bg-[#fa4616]');
        // repalce with dropColor
        e.target.classList.remove(dropColor);
    }}
>
    <slot>Drag & Drop a file here</slot>

    <AlertDialog.Root bind:open={showPremiumDialog}>
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
                <AlertDialog.Action>Awesome</AlertDialog.Action>
            </AlertDialog.Footer>
        </AlertDialog.Content>
    </AlertDialog.Root>
</article>
