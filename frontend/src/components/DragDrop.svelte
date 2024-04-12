<script lang="ts">
    import ShowPremiumDialog from '$components/Premium.svelte';

    export let fileTypes: string[] = [],
        dropColor: string = 'bg-stone-400',
        exportType: string = 'storyline',
        disabled: boolean = false,
        onDrop: (files: FileList | null, resultText: string, exportType: string) => void;

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
                    `<span class="text-2xl">😧</span> Seriously? Please only use ${fileTypes.slice(0, -1).join(', ')} and ${fileTypes.slice(-1)} files!`,
                    exportType
                );
                return;
            }
        }

        onDrop(
            files,
            `Processing ${files.length > 1 ? files.length + ' files!' : '<span class="text-cyan-400">' + files[0].name + '</span>'}`,
            exportType
        );
    }
</script>

<article
    id="drop"
    class="border-2 border-dotted p-4 rounded-lg bg-opacity-60 max-w-60"
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

    <ShowPremiumDialog show={showPremiumDialog} />
</article>
