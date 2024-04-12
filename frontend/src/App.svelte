<script lang="ts">
    import './app.css';

    import excel from '$assets/images/excel.svg';
    import word from '$assets/images/word.svg';
    import csv from '$assets/images/csv.svg';

    import { Quiz } from '../wailsjs/go/main/App.js';

    import Header from '$components/Header.svelte';
    import DragDrop from '$components/DragDrop.svelte';

    import { Progress } from '$lib/components/ui/progress';
    import { Switch } from '$lib/components/ui/switch';

    let progress = 0;

    let message: string =
        '<span class="text-2xl hidden xs:inline">👇</span> Please drop your files here <span class="text-2xl">👇</span>';
    let fileName: string;

    async function onDrop(files: FileList | null, resultText: string, outputType: string): Promise<void> {
        message = resultText;

        let fileIndex = 0;

        if (files && files.length > 0) {
            const progressInterval = setInterval(() => {
                progress++;
                console.log('progress:', progress);
                if (progress >= 100) {
                    setTimeout(() => {
                        clearInterval(progressInterval);
                        progress = 0;
                    }, 200);
                }
            }, 20);

            for (const file of files) {
                fileName = file.name;

                const fileType = fileName.split('.').pop() as string; // file extension

                console.log('fileType:', fileType);

                const fileContents = await readFileAsBase64(file);
                message = await Quiz(fileContents, fileName, fileType, outputType);

                fileIndex = [...files].indexOf(file);
                progress = ((fileIndex + 1) / files.length) * 100;
            }

            if (files.length > 1) {
                message = `<span class="text-2xl">🎉</span> ${files.length} files exported! <span class="text-2xl">🚀</span>`;
            }
        }
    }

    async function readFileAsBase64(file: File): Promise<string> {
        return new Promise((resolve, reject) => {
            const reader = new FileReader();

            reader.onload = () => {
                resolve(btoa(reader.result as string));
            };

            reader.onerror = reject;
            reader.readAsBinaryString(file);
        });
    }
</script>

<main class="flex flex-col items-center overflow-auto">
    <h1 class="text-4xl font-black mt-4 text-orange-600">StoryQuiz</h1>
    <h2 class="text-sm">(Quiz Exporter)</h2>
    <Header />
    <Switch class="hidden" />
    <div id="message" class="xs:leading-10">{@html message}</div>
    <section class="dragdrop flex justify-center flex-wrap gap-4 mt-2 xs:mt-0">
        <DragDrop fileTypes={['docx', 'xslx', 'xlsm']} dropColor={'bg-storyline'} {onDrop}>
            Drop a Word or Excel file here for Storyline Quiz
        </DragDrop>
        <DragDrop fileTypes={['docx', 'xlsx', 'xlsm']} dropColor={'bg-csod'} outputType={'csod'} {onDrop}>
            Drop a Word or Excel file here for Cornerstone Quiz
        </DragDrop>
    </section>
    <footer class="w-full p-4">
        <Progress value={progress} max={100} class={`max-w-60 mx-auto ${progress ? 'visible' : 'invisible'}`} />
    </footer>
</main>
