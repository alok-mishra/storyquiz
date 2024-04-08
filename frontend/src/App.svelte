<script lang="ts">
    // import logo from './assets/images/logo-universal.png';

    import storyquiz from './assets/images/storyquiz.svg';
    import storyline from './assets/images/storyline.svg';
    import csod from './assets/images/csod.svg';
    import excel from './assets/images/excel.svg';
    import word from './assets/images/word.svg';
    import csv from './assets/images/csv.svg';
    import './app.css';
    import { Quiz } from '../wailsjs/go/main/App.js';

    import DragDrop from './components/DragDrop.svelte';

    import { Progress } from '$lib/components/ui/progress';
    import { Switch } from '$lib/components/ui/switch';
    import * as Popover from '$lib/components/ui/popover';
    import * as Drawer from '$lib/components/ui/drawer';
    import { Button } from '$lib/components/ui/button';

    let progress = 0;

    let message: string = '👇 Please drop your file here 👇';
    let fileName: string;

    function onDrop(file: File | null, resultText: string): void {
        message = resultText;

        if (file) {
            const progressInterval = setInterval(() => {
                progress++;
                console.log('progress:', progress);
                if (progress >= 100) {
                    setTimeout(() => {
                        progress = 0;
                        clearInterval(progressInterval);
                    }, 600);
                }
            }, 10);

            fileName = file.name;
            const fileType = fileName.split('.').pop() as string; // file extension

            const reader = new FileReader();
            reader.onload = async () => {
                const fileContents = reader.result as string;
                await Quiz(btoa(fileContents), fileName, fileType).then((result) => {
                    progress = 100;
                    message = result;
                });
            };
            reader.readAsBinaryString(file);
        }
    }
</script>

<main class="flex flex-col items-center overflow-auto">
    <h1 class="text-4xl font-black mt-4 text-orange-600">StoryQuiz</h1>
    <h2 class="text-sm">(Quiz Exporter)</h2>
    <section class="relative">
        <Popover.Root>
            <Popover.Trigger>
                <img class="w-60 p-4" alt="StoryQuiz Logo" src={storyquiz} />
            </Popover.Trigger>
            <Popover.Content class="w-60 !top-12">
                <div class="flex justify-around items-center">
                    <div>
                        <img class="w-20 p-4" alt="StoryQuiz Logo" src={storyquiz} />
                    </div>
                    <div>
                        <p class="text-lg font-bold text-orange-600">Story Quiz</p>
                        <h4 class="text-sm font-semibold text-cyan-200">@alok.mishra</h4>
                        <span class="text-xs text-muted-foreground">v1.0, April 2024</span>
                        <Drawer.Root>
                            <Drawer.Trigger class="text-sm text-cyan-600 hover:text-cyan-400">License</Drawer.Trigger>
                            <Drawer.Content class="mx-10">
                                <Drawer.Header>
                                    <Drawer.Title class="border-b border-solid mb-4 pb-2">MIT License</Drawer.Title>
                                    <Drawer.Description>
                                        <section class="flex flex-col gap-4">
                                            <p>Copyright 2024 Alok Mishra</p>
                                            <p>
                                                Permission is hereby granted, free of charge, to any person obtaining a
                                                copy of this software and associated documentation files (the
                                                “Software”), to deal in the Software without restriction, including
                                                without limitation the rights to use, copy, modify, merge, publish,
                                                distribute, sublicense, and/or sell copies of the Software, and to
                                                permit persons to whom the Software is furnished to do so, subject to
                                                the following conditions:
                                            </p>
                                            <p>
                                                The above copyright notice and this permission notice shall be included
                                                in all copies or substantial portions of the Software.
                                            </p>
                                            <p>
                                                THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS
                                                OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
                                                MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
                                                IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
                                                CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
                                                TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
                                                SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
                                            </p>
                                        </section>
                                    </Drawer.Description>
                                </Drawer.Header>
                                <Drawer.Footer>
                                    <Drawer.Close><Button class="font-bold">Close</Button></Drawer.Close>
                                </Drawer.Footer>
                            </Drawer.Content>
                        </Drawer.Root>
                    </div>
                </div>
            </Popover.Content>
        </Popover.Root>
        <img class="w-40 p-4 absolute -bottom-0 -left-20" alt="Storyline Logo" src={storyline} />
        <img class="w-40 p-4 absolute -bottom-0 -right-20" alt="CSOD Logo" src={csod} />
    </section>
    <Switch class="bg-orange-400" />
    <div id="message">{@html message}</div>
    <section class="dragdrop flex">
        <DragDrop dropColor={'bg-storyline'} fileTypes={['doc', 'docx', 'xls', 'xslx', 'xlsm']} {onDrop}>
            Drop a Word or Excel file here for Storyline Quiz
        </DragDrop>
        <DragDrop dropColor={'bg-csod'} fileTypes={['xls', 'xlsx', 'xlsm']} {onDrop} disabled>
            Drop a Excel file here for CSOD Quiz
        </DragDrop>
    </section>
    <footer class="w-full p-4">
        <Progress value={progress} max={100} class={`max-w-60 mx-auto ${progress ? 'visible' : 'invisible'}`} />
    </footer>
</main>
