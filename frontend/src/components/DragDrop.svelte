<script lang="ts">
    import * as AlertDialog from '$lib/components/ui/alert-dialog/index.js';
    export let fileTypes: string[] = [],
        onDrop: (files: FileList | null, resultText: string) => void,
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

        const files = event.dataTransfer?.files;

        if (!files || files.length === 0) {
            onDrop(null, 'No files dropped!');
            return;
        }

        for (const file of files) {
            const fileType = file.name.split('.').pop() as string;

            if (!fileTypes.includes(fileType)) {
                onDrop(
                    null,
                    `<span class="text-2xl">😧</span> Seriously? Please only use ${fileTypes.slice(0, -1).join(', ')} and ${fileTypes.slice(-1)} files!`
                );
                return;
            }
        }

        onDrop(
            files,
            `Processing ${files.length > 1 ? files.length + ' files!' : '<span class="text-cyan-400">' + files[0].name + '</span>'}`
        );
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
                <AlertDialog.Action class="font-bold">Awesome</AlertDialog.Action>
            </AlertDialog.Footer>
        </AlertDialog.Content>
    </AlertDialog.Root>
</article>
