package output

import (
	"dbcompare/compare"
	"dbcompare/stats"
	"html/template"
	"os"

	"github.com/fatih/color"
)

type DadosOutput struct {
	Resultado       compare.ResultadoGeral
	BancoPrimario   string
	BancoSecundario string
	Version         string
	Stats           stats.Stats
}

func GerarHtml(dadosOutput DadosOutput) error {

	strTemplate := `<!DOCTYPE html>
	<html lang="pt-br">
	
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
	
		<title>DB Compare</title>
	</head>
	
	<style>
	/*! tailwindcss v3.4.3 | MIT License | https://tailwindcss.com*/*,:after,:before{box-sizing:border-box;border:0 solid #e5e7eb}:after,:before{--tw-content:""}:host,html{line-height:1.5;-webkit-text-size-adjust:100%;-moz-tab-size:4;-o-tab-size:4;tab-size:4;font-family:ui-sans-serif,system-ui,sans-serif,Apple Color Emoji,Segoe UI Emoji,Segoe UI Symbol,Noto Color Emoji;font-feature-settings:normal;font-variation-settings:normal;-webkit-tap-highlight-color:transparent}body{margin:0;line-height:inherit}hr{height:0;color:inherit;border-top-width:1px}abbr:where([title]){-webkit-text-decoration:underline dotted;text-decoration:underline dotted}h1,h2,h3,h4,h5,h6{font-size:inherit;font-weight:inherit}a{color:inherit;text-decoration:inherit}b,strong{font-weight:bolder}code,kbd,pre,samp{font-family:ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,Liberation Mono,Courier New,monospace;font-feature-settings:normal;font-variation-settings:normal;font-size:1em}small{font-size:80%}sub,sup{font-size:75%;line-height:0;position:relative;vertical-align:initial}sub{bottom:-.25em}sup{top:-.5em}table{text-indent:0;border-color:inherit;border-collapse:collapse}button,input,optgroup,select,textarea{font-family:inherit;font-feature-settings:inherit;font-variation-settings:inherit;font-size:100%;font-weight:inherit;line-height:inherit;letter-spacing:inherit;color:inherit;margin:0;padding:0}button,select{text-transform:none}button,input:where([type=button]),input:where([type=reset]),input:where([type=submit]){-webkit-appearance:button;background-color:initial;background-image:none}:-moz-focusring{outline:auto}:-moz-ui-invalid{box-shadow:none}progress{vertical-align:initial}::-webkit-inner-spin-button,::-webkit-outer-spin-button{height:auto}[type=search]{-webkit-appearance:textfield;outline-offset:-2px}::-webkit-search-decoration{-webkit-appearance:none}::-webkit-file-upload-button{-webkit-appearance:button;font:inherit}summary{display:list-item}blockquote,dd,dl,figure,h1,h2,h3,h4,h5,h6,hr,p,pre{margin:0}fieldset{margin:0}fieldset,legend{padding:0}menu,ol,ul{list-style:none;margin:0;padding:0}dialog{padding:0}textarea{resize:vertical}input::-moz-placeholder,textarea::-moz-placeholder{opacity:1;color:#9ca3af}input::placeholder,textarea::placeholder{opacity:1;color:#9ca3af}[role=button],button{cursor:pointer}:disabled{cursor:default}audio,canvas,embed,iframe,img,object,svg,video{display:block;vertical-align:middle}img,video{max-width:100%;height:auto}[hidden]{display:none}*,::backdrop,:after,:before{--tw-border-spacing-x:0;--tw-border-spacing-y:0;--tw-translate-x:0;--tw-translate-y:0;--tw-rotate:0;--tw-skew-x:0;--tw-skew-y:0;--tw-scale-x:1;--tw-scale-y:1;--tw-pan-x: ;--tw-pan-y: ;--tw-pinch-zoom: ;--tw-scroll-snap-strictness:proximity;--tw-gradient-from-position: ;--tw-gradient-via-position: ;--tw-gradient-to-position: ;--tw-ordinal: ;--tw-slashed-zero: ;--tw-numeric-figure: ;--tw-numeric-spacing: ;--tw-numeric-fraction: ;--tw-ring-inset: ;--tw-ring-offset-width:0px;--tw-ring-offset-color:#fff;--tw-ring-color:#3b82f680;--tw-ring-offset-shadow:0 0 #0000;--tw-ring-shadow:0 0 #0000;--tw-shadow:0 0 #0000;--tw-shadow-colored:0 0 #0000;--tw-blur: ;--tw-brightness: ;--tw-contrast: ;--tw-grayscale: ;--tw-hue-rotate: ;--tw-invert: ;--tw-saturate: ;--tw-sepia: ;--tw-drop-shadow: ;--tw-backdrop-blur: ;--tw-backdrop-brightness: ;--tw-backdrop-contrast: ;--tw-backdrop-grayscale: ;--tw-backdrop-hue-rotate: ;--tw-backdrop-invert: ;--tw-backdrop-opacity: ;--tw-backdrop-saturate: ;--tw-backdrop-sepia: ;--tw-contain-size: ;--tw-contain-layout: ;--tw-contain-paint: ;--tw-contain-style: }.container{width:100%}@media (min-width:640px){.container{max-width:640px}}@media (min-width:768px){.container{max-width:768px}}@media (min-width:1024px){.container{max-width:1024px}}@media (min-width:1280px){.container{max-width:1280px}}@media (min-width:1536px){.container{max-width:1536px}}.sr-only{position:absolute;width:1px;height:1px;padding:0;margin:-1px;overflow:hidden;clip:rect(0,0,0,0);white-space:nowrap;border-width:0}.not-sr-only{position:static;width:auto;height:auto;padding:0;margin:0;overflow:visible;clip:auto;white-space:normal}.pointer-events-none{pointer-events:none}.pointer-events-auto{pointer-events:auto}.\!visible{visibility:visible!important}.visible{visibility:visible}.invisible{visibility:hidden}.collapse{visibility:collapse}.static{position:static}.fixed{position:fixed}.absolute{position:absolute}.relative{position:relative}.sticky{position:sticky}.-inset-1{inset:-.25rem}.end-1{inset-inline-end:.25rem}.isolate{isolation:isolate}.isolation-auto{isolation:auto}.float-start{float:inline-start}.float-end{float:inline-end}.float-right{float:right}.float-left{float:left}.float-none{float:none}.clear-start{clear:inline-start}.clear-end{clear:inline-end}.clear-left{clear:left}.clear-right{clear:right}.clear-both{clear:both}.clear-none{clear:none}.m-10{margin:2.5rem}.ml-5{margin-left:1.25rem}.mt-2{margin-top:.5rem}.box-border{box-sizing:border-box}.box-content{box-sizing:initial}.line-clamp-none{overflow:visible;display:block;-webkit-box-orient:horizontal;-webkit-line-clamp:none}.block{display:block}.inline-block{display:inline-block}.inline{display:inline}.flex{display:flex}.inline-flex{display:inline-flex}.table{display:table}.inline-table{display:inline-table}.table-caption{display:table-caption}.table-cell{display:table-cell}.table-column{display:table-column}.table-column-group{display:table-column-group}.table-footer-group{display:table-footer-group}.table-header-group{display:table-header-group}.table-row-group{display:table-row-group}.table-row{display:table-row}.flow-root{display:flow-root}.grid{display:grid}.inline-grid{display:inline-grid}.contents{display:contents}.list-item{display:list-item}.hidden{display:none}.w-10{width:2.5rem}.w-10\/12{width:83.333333%}.w-96{width:24rem}.w-\[this-is\\\\\]{width:this-is\\}.w-\[this-is\]{width:this-is}.w-\[weird-and-invalid\]{width:weird-and-invalid}.w-full{width:100%}.w-1\/2{width:50%}.min-w-full{min-width:100%}.flex-1{flex:1 1 0%}.flex-shrink,.shrink{flex-shrink:1}.flex-grow,.grow{flex-grow:1}.table-auto{table-layout:auto}.table-fixed{table-layout:fixed}.caption-top{caption-side:top}.caption-bottom{caption-side:bottom}.border-collapse{border-collapse:collapse}.border-separate{border-collapse:initial}.border-spacing-3{--tw-border-spacing-x:0.75rem;--tw-border-spacing-y:0.75rem;border-spacing:var(--tw-border-spacing-x) var(--tw-border-spacing-y)}.\!transform{transform:translate(var(--tw-translate-x),var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))!important}.transform,.transform-cpu{transform:translate(var(--tw-translate-x),var(--tw-translate-y)) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))}.transform-gpu{transform:translate3d(var(--tw-translate-x),var(--tw-translate-y),0) rotate(var(--tw-rotate)) skewX(var(--tw-skew-x)) skewY(var(--tw-skew-y)) scaleX(var(--tw-scale-x)) scaleY(var(--tw-scale-y))}.transform-none{transform:none}.touch-auto{touch-action:auto}.touch-none{touch-action:none}.touch-pan-x{--tw-pan-x:pan-x}.touch-pan-left,.touch-pan-x{touch-action:var(--tw-pan-x) var(--tw-pan-y) var(--tw-pinch-zoom)}.touch-pan-left{--tw-pan-x:pan-left}.touch-pan-right{--tw-pan-x:pan-right}.touch-pan-right,.touch-pan-y{touch-action:var(--tw-pan-x) var(--tw-pan-y) var(--tw-pinch-zoom)}.touch-pan-y{--tw-pan-y:pan-y}.touch-pan-up{--tw-pan-y:pan-up}.touch-pan-down,.touch-pan-up{touch-action:var(--tw-pan-x) var(--tw-pan-y) var(--tw-pinch-zoom)}.touch-pan-down{--tw-pan-y:pan-down}.touch-pinch-zoom{--tw-pinch-zoom:pinch-zoom;touch-action:var(--tw-pan-x) var(--tw-pan-y) var(--tw-pinch-zoom)}.touch-manipulation{touch-action:manipulation}.select-none{-webkit-user-select:none;-moz-user-select:none;user-select:none}.select-text{-webkit-user-select:text;-moz-user-select:text;user-select:text}.select-all{-webkit-user-select:all;-moz-user-select:all;user-select:all}.select-auto{-webkit-user-select:auto;-moz-user-select:auto;user-select:auto}.resize-none{resize:none}.resize-y{resize:vertical}.resize-x{resize:horizontal}.resize{resize:both}.snap-none{scroll-snap-type:none}.snap-x{scroll-snap-type:x var(--tw-scroll-snap-strictness)}.snap-y{scroll-snap-type:y var(--tw-scroll-snap-strictness)}.snap-both{scroll-snap-type:both var(--tw-scroll-snap-strictness)}.snap-mandatory{--tw-scroll-snap-strictness:mandatory}.snap-proximity{--tw-scroll-snap-strictness:proximity}.snap-start{scroll-snap-align:start}.snap-end{scroll-snap-align:end}.snap-center{scroll-snap-align:center}.snap-align-none{scroll-snap-align:none}.snap-normal{scroll-snap-stop:normal}.snap-always{scroll-snap-stop:always}.list-inside{list-style-position:inside}.list-outside{list-style-position:outside}.appearance-none{-webkit-appearance:none;-moz-appearance:none;appearance:none}.appearance-auto{-webkit-appearance:auto;-moz-appearance:auto;appearance:auto}.break-before-auto{-moz-column-break-before:auto;break-before:auto}.break-before-avoid{-moz-column-break-before:avoid;break-before:avoid}.break-before-all{-moz-column-break-before:all;break-before:all}.break-before-avoid-page{-moz-column-break-before:avoid;break-before:avoid-page}.break-before-page{-moz-column-break-before:page;break-before:page}.break-before-left{-moz-column-break-before:left;break-before:left}.break-before-right{-moz-column-break-before:right;break-before:right}.break-before-column{-moz-column-break-before:column;break-before:column}.break-inside-auto{-moz-column-break-inside:auto;break-inside:auto}.break-inside-avoid{-moz-column-break-inside:avoid;break-inside:avoid}.break-inside-avoid-page{break-inside:avoid-page}.break-inside-avoid-column{-moz-column-break-inside:avoid;break-inside:avoid-column}.break-after-auto{-moz-column-break-after:auto;break-after:auto}.break-after-avoid{-moz-column-break-after:avoid;break-after:avoid}.break-after-all{-moz-column-break-after:all;break-after:all}.break-after-avoid-page{-moz-column-break-after:avoid;break-after:avoid-page}.break-after-page{-moz-column-break-after:page;break-after:page}.break-after-left{-moz-column-break-after:left;break-after:left}.break-after-right{-moz-column-break-after:right;break-after:right}.break-after-column{-moz-column-break-after:column;break-after:column}.grid-flow-row{grid-auto-flow:row}.grid-flow-col{grid-auto-flow:column}.grid-flow-dense{grid-auto-flow:dense}.grid-flow-row-dense{grid-auto-flow:row dense}.grid-flow-col-dense{grid-auto-flow:column dense}.flex-row{flex-direction:row}.flex-row-reverse{flex-direction:row-reverse}.flex-col{flex-direction:column}.flex-col-reverse{flex-direction:column-reverse}.flex-wrap{flex-wrap:wrap}.flex-wrap-reverse{flex-wrap:wrap-reverse}.flex-nowrap{flex-wrap:nowrap}.place-content-center{place-content:center}.place-content-start{place-content:start}.place-content-end{place-content:end}.place-content-between{place-content:space-between}.place-content-around{place-content:space-around}.place-content-evenly{place-content:space-evenly}.place-content-baseline{place-content:baseline}.place-content-stretch{place-content:stretch}.place-items-start{place-items:start}.place-items-end{place-items:end}.place-items-center{place-items:center}.place-items-baseline{place-items:baseline}.place-items-stretch{place-items:stretch}.content-normal{align-content:normal}.content-center{align-content:center}.content-start{align-content:flex-start}.content-end{align-content:flex-end}.content-between{align-content:space-between}.content-around{align-content:space-around}.content-evenly{align-content:space-evenly}.content-baseline{align-content:baseline}.content-stretch{align-content:stretch}.items-start{align-items:flex-start}.items-end{align-items:flex-end}.items-center{align-items:center}.items-baseline{align-items:baseline}.items-stretch{align-items:stretch}.justify-normal{justify-content:normal}.justify-start{justify-content:flex-start}.justify-end{justify-content:flex-end}.justify-center{justify-content:center}.justify-between{justify-content:space-between}.justify-around{justify-content:space-around}.justify-evenly{justify-content:space-evenly}.justify-stretch{justify-content:stretch}.justify-items-start{justify-items:start}.justify-items-end{justify-items:end}.justify-items-center{justify-items:center}.justify-items-stretch{justify-items:stretch}.gap-10{gap:2.5rem}.space-x-2>:not([hidden])~:not([hidden]){--tw-space-x-reverse:0;margin-right:calc(.5rem*var(--tw-space-x-reverse));margin-left:calc(.5rem*(1 - var(--tw-space-x-reverse)))}.space-y-reverse>:not([hidden])~:not([hidden]){--tw-space-y-reverse:1}.space-x-reverse>:not([hidden])~:not([hidden]){--tw-space-x-reverse:1}.divide-x>:not([hidden])~:not([hidden]){--tw-divide-x-reverse:0;border-right-width:calc(1px*var(--tw-divide-x-reverse));border-left-width:calc(1px*(1 - var(--tw-divide-x-reverse)))}.divide-y>:not([hidden])~:not([hidden]){--tw-divide-y-reverse:0;border-top-width:calc(1px*(1 - var(--tw-divide-y-reverse)));border-bottom-width:calc(1px*var(--tw-divide-y-reverse))}.divide-y-reverse>:not([hidden])~:not([hidden]){--tw-divide-y-reverse:1}.divide-x-reverse>:not([hidden])~:not([hidden]){--tw-divide-x-reverse:1}.divide-solid>:not([hidden])~:not([hidden]){border-style:solid}.divide-dashed>:not([hidden])~:not([hidden]){border-style:dashed}.divide-dotted>:not([hidden])~:not([hidden]){border-style:dotted}.divide-double>:not([hidden])~:not([hidden]){border-style:double}.divide-none>:not([hidden])~:not([hidden]){border-style:none}.place-self-auto{place-self:auto}.place-self-start{place-self:start}.place-self-end{place-self:end}.place-self-center{place-self:center}.place-self-stretch{place-self:stretch}.self-auto{align-self:auto}.self-start{align-self:flex-start}.self-end{align-self:flex-end}.self-center{align-self:center}.self-stretch{align-self:stretch}.self-baseline{align-self:baseline}.justify-self-auto{justify-self:auto}.justify-self-start{justify-self:start}.justify-self-end{justify-self:end}.justify-self-center{justify-self:center}.justify-self-stretch{justify-self:stretch}.overflow-auto{overflow:auto}.overflow-hidden{overflow:hidden}.overflow-clip{overflow:clip}.overflow-visible{overflow:visible}.overflow-scroll{overflow:scroll}.overflow-x-auto{overflow-x:auto}.overflow-y-auto{overflow-y:auto}.overflow-x-hidden{overflow-x:hidden}.overflow-y-hidden{overflow-y:hidden}.overflow-x-clip{overflow-x:clip}.overflow-y-clip{overflow-y:clip}.overflow-x-visible{overflow-x:visible}.overflow-y-visible{overflow-y:visible}.overflow-x-scroll{overflow-x:scroll}.overflow-y-scroll{overflow-y:scroll}.overscroll-auto{overscroll-behavior:auto}.overscroll-contain{overscroll-behavior:contain}.overscroll-none{overscroll-behavior:none}.overscroll-y-auto{overscroll-behavior-y:auto}.overscroll-y-contain{overscroll-behavior-y:contain}.overscroll-y-none{overscroll-behavior-y:none}.overscroll-x-auto{overscroll-behavior-x:auto}.overscroll-x-contain{overscroll-behavior-x:contain}.overscroll-x-none{overscroll-behavior-x:none}.scroll-auto{scroll-behavior:auto}.scroll-smooth{scroll-behavior:smooth}.truncate{overflow:hidden;white-space:nowrap}.overflow-ellipsis,.text-ellipsis,.truncate{text-overflow:ellipsis}.text-clip{text-overflow:clip}.hyphens-none{-webkit-hyphens:none;hyphens:none}.hyphens-manual{-webkit-hyphens:manual;hyphens:manual}.hyphens-auto{-webkit-hyphens:auto;hyphens:auto}.whitespace-normal{white-space:normal}.whitespace-nowrap{white-space:nowrap}.whitespace-pre{white-space:pre}.whitespace-pre-line{white-space:pre-line}.whitespace-pre-wrap{white-space:pre-wrap}.whitespace-break-spaces{white-space:break-spaces}.text-wrap{text-wrap:wrap}.text-nowrap{text-wrap:nowrap}.text-balance{text-wrap:balance}.text-pretty{text-wrap:pretty}.break-normal{overflow-wrap:normal;word-break:normal}.break-words{overflow-wrap:break-word}.break-all{word-break:break-all}.break-keep{word-break:keep-all}.rounded{border-radius:.25rem}.rounded-lg{border-radius:.5rem}.rounded-md{border-radius:.375rem}.rounded-b{border-bottom-right-radius:.25rem;border-bottom-left-radius:.25rem}.rounded-e{border-start-end-radius:.25rem;border-end-end-radius:.25rem}.rounded-l{border-top-left-radius:.25rem;border-bottom-left-radius:.25rem}.rounded-r{border-top-right-radius:.25rem;border-bottom-right-radius:.25rem}.rounded-s{border-start-start-radius:.25rem;border-end-start-radius:.25rem}.rounded-t{border-top-left-radius:.25rem;border-top-right-radius:.25rem}.rounded-bl{border-bottom-left-radius:.25rem}.rounded-br{border-bottom-right-radius:.25rem}.rounded-ee{border-end-end-radius:.25rem}.rounded-es{border-end-start-radius:.25rem}.rounded-se{border-start-end-radius:.25rem}.rounded-ss{border-start-start-radius:.25rem}.rounded-tl{border-top-left-radius:.25rem}.rounded-tr{border-top-right-radius:.25rem}.border{border-width:1px}.border-2{border-width:2px}.border-x{border-left-width:1px;border-right-width:1px}.border-y{border-top-width:1px}.border-b,.border-y{border-bottom-width:1px}.border-b-2{border-bottom-width:2px}.border-e{border-inline-end-width:1px}.border-l{border-left-width:1px}.border-r{border-right-width:1px}.border-s{border-inline-start-width:1px}.border-t{border-top-width:1px}.border-solid{border-style:solid}.border-dashed{border-style:dashed}.border-dotted{border-style:dotted}.border-double{border-style:double}.border-hidden{border-style:hidden}.border-none{border-style:none}.border-b-black{--tw-border-opacity:1;border-bottom-color:rgb(0 0 0/var(--tw-border-opacity))}.border-b-slate-400{--tw-border-opacity:1;border-bottom-color:rgb(148 163 184/var(--tw-border-opacity))}.bg-\[rgb\(255\2c 0\2c 0\)\]{--tw-bg-opacity:1;background-color:rgb(255 0 0/var(--tw-bg-opacity))}.bg-gray-200{--tw-bg-opacity:1;background-color:rgb(229 231 235/var(--tw-bg-opacity))}.bg-green-200{--tw-bg-opacity:1;background-color:rgb(187 247 208/var(--tw-bg-opacity))}.bg-red-200{--tw-bg-opacity:1;background-color:rgb(254 202 202/var(--tw-bg-opacity))}.bg-slate-500{--tw-bg-opacity:1;background-color:rgb(100 116 139/var(--tw-bg-opacity))}.bg-slate-800{--tw-bg-opacity:1;background-color:rgb(30 41 59/var(--tw-bg-opacity))}.bg-white{--tw-bg-opacity:1;background-color:rgb(255 255 255/var(--tw-bg-opacity))}.decoration-slice{-webkit-box-decoration-break:slice;box-decoration-break:slice}.decoration-clone{-webkit-box-decoration-break:clone;box-decoration-break:clone}.box-decoration-slice{-webkit-box-decoration-break:slice;box-decoration-break:slice}.box-decoration-clone{-webkit-box-decoration-break:clone;box-decoration-break:clone}.bg-fixed{background-attachment:fixed}.bg-local{background-attachment:local}.bg-scroll{background-attachment:scroll}.bg-clip-border{background-clip:initial}.bg-clip-padding{background-clip:padding-box}.bg-clip-content{background-clip:content-box}.bg-clip-text{-webkit-background-clip:text;background-clip:text}.bg-repeat{background-repeat:repeat}.bg-no-repeat{background-repeat:no-repeat}.bg-repeat-x{background-repeat:repeat-x}.bg-repeat-y{background-repeat:repeat-y}.bg-repeat-round{background-repeat:round}.bg-repeat-space{background-repeat:space}.bg-origin-border{background-origin:border-box}.bg-origin-padding{background-origin:initial}.bg-origin-content{background-origin:content-box}.object-contain{-o-object-fit:contain;object-fit:contain}.object-cover{-o-object-fit:cover;object-fit:cover}.object-fill{-o-object-fit:fill;object-fit:fill}.object-none{-o-object-fit:none;object-fit:none}.object-scale-down{-o-object-fit:scale-down;object-fit:scale-down}.p-4{padding:1rem}.p-6{padding:1.5rem}.p-7{padding:1.75rem}.p-10{padding:2.5rem}.p-3{padding:.75rem}.px-4{padding-left:1rem;padding-right:1rem}.py-2{padding-top:.5rem}.pb-2,.py-2{padding-bottom:.5rem}.pt-10{padding-top:2.5rem}.pt-4{padding-top:1rem}.text-left{text-align:left}.text-center{text-align:center}.text-right{text-align:right}.text-justify{text-align:justify}.text-start{text-align:start}.text-end{text-align:end}.align-baseline{vertical-align:initial}.align-top{vertical-align:top}.align-middle{vertical-align:middle}.align-bottom{vertical-align:bottom}.align-text-top{vertical-align:text-top}.align-text-bottom{vertical-align:text-bottom}.align-sub{vertical-align:sub}.align-super{vertical-align:super}.font-mono{font-family:ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,Liberation Mono,Courier New,monospace}.text-2xl{font-size:1.5rem;line-height:2rem}.text-3xl{font-size:1.875rem;line-height:2.25rem}.text-xl{font-size:1.25rem;line-height:1.75rem}.font-bold{font-weight:700}.font-semibold{font-weight:600}.uppercase{text-transform:uppercase}.lowercase{text-transform:lowercase}.capitalize{text-transform:capitalize}.normal-case{text-transform:none}.italic{font-style:italic}.not-italic{font-style:normal}.normal-nums{font-variant-numeric:normal}.ordinal{--tw-ordinal:ordinal}.ordinal,.slashed-zero{font-variant-numeric:var(--tw-ordinal) var(--tw-slashed-zero) var(--tw-numeric-figure) var(--tw-numeric-spacing) var(--tw-numeric-fraction)}.slashed-zero{--tw-slashed-zero:slashed-zero}.lining-nums{--tw-numeric-figure:lining-nums}.lining-nums,.oldstyle-nums{font-variant-numeric:var(--tw-ordinal) var(--tw-slashed-zero) var(--tw-numeric-figure) var(--tw-numeric-spacing) var(--tw-numeric-fraction)}.oldstyle-nums{--tw-numeric-figure:oldstyle-nums}.proportional-nums{--tw-numeric-spacing:proportional-nums}.proportional-nums,.tabular-nums{font-variant-numeric:var(--tw-ordinal) var(--tw-slashed-zero) var(--tw-numeric-figure) var(--tw-numeric-spacing) var(--tw-numeric-fraction)}.tabular-nums{--tw-numeric-spacing:tabular-nums}.diagonal-fractions{--tw-numeric-fraction:diagonal-fractions}.diagonal-fractions,.stacked-fractions{font-variant-numeric:var(--tw-ordinal) var(--tw-slashed-zero) var(--tw-numeric-figure) var(--tw-numeric-spacing) var(--tw-numeric-fraction)}.stacked-fractions{--tw-numeric-fraction:stacked-fractions}.text-\[\#336699\]\/\[\.35\]{color:#33669959}.text-blue-500{--tw-text-opacity:1;color:rgb(59 130 246/var(--tw-text-opacity))}.text-green-500{--tw-text-opacity:1;color:rgb(34 197 94/var(--tw-text-opacity))}.text-red-600{--tw-text-opacity:1;color:rgb(220 38 38/var(--tw-text-opacity))}.text-white{--tw-text-opacity:1;color:rgb(255 255 255/var(--tw-text-opacity))}.underline{text-decoration-line:underline}.overline{text-decoration-line:overline}.line-through{text-decoration-line:line-through}.no-underline{text-decoration-line:none}.decoration-solid{text-decoration-style:solid}.decoration-double{text-decoration-style:double}.decoration-dotted{text-decoration-style:dotted}.decoration-dashed{text-decoration-style:dashed}.decoration-wavy{text-decoration-style:wavy}.antialiased{-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale}.subpixel-antialiased{-webkit-font-smoothing:auto;-moz-osx-font-smoothing:auto}.bg-blend-normal{background-blend-mode:normal}.bg-blend-multiply{background-blend-mode:multiply}.bg-blend-screen{background-blend-mode:screen}.bg-blend-overlay{background-blend-mode:overlay}.bg-blend-darken{background-blend-mode:darken}.bg-blend-lighten{background-blend-mode:lighten}.bg-blend-color-dodge{background-blend-mode:color-dodge}.bg-blend-color-burn{background-blend-mode:color-burn}.bg-blend-hard-light{background-blend-mode:hard-light}.bg-blend-soft-light{background-blend-mode:soft-light}.bg-blend-difference{background-blend-mode:difference}.bg-blend-exclusion{background-blend-mode:exclusion}.bg-blend-hue{background-blend-mode:hue}.bg-blend-saturation{background-blend-mode:saturation}.bg-blend-color{background-blend-mode:color}.bg-blend-luminosity{background-blend-mode:luminosity}.mix-blend-normal{mix-blend-mode:normal}.mix-blend-multiply{mix-blend-mode:multiply}.mix-blend-screen{mix-blend-mode:screen}.mix-blend-overlay{mix-blend-mode:overlay}.mix-blend-darken{mix-blend-mode:darken}.mix-blend-lighten{mix-blend-mode:lighten}.mix-blend-color-dodge{mix-blend-mode:color-dodge}.mix-blend-color-burn{mix-blend-mode:color-burn}.mix-blend-hard-light{mix-blend-mode:hard-light}.mix-blend-soft-light{mix-blend-mode:soft-light}.mix-blend-difference{mix-blend-mode:difference}.mix-blend-exclusion{mix-blend-mode:exclusion}.mix-blend-hue{mix-blend-mode:hue}.mix-blend-saturation{mix-blend-mode:saturation}.mix-blend-color{mix-blend-mode:color}.mix-blend-luminosity{mix-blend-mode:luminosity}.mix-blend-plus-darker{mix-blend-mode:plus-darker}.mix-blend-plus-lighter{mix-blend-mode:plus-lighter}.\!shadow{--tw-shadow:0 1px 3px 0 #0000001a,0 1px 2px -1px #0000001a!important;--tw-shadow-colored:0 1px 3px 0 var(--tw-shadow-color),0 1px 2px -1px var(--tw-shadow-color)!important;box-shadow:var(--tw-ring-offset-shadow,0 0 #0000),var(--tw-ring-shadow,0 0 #0000),var(--tw-shadow)!important}.shadow{--tw-shadow:0 1px 3px 0 #0000001a,0 1px 2px -1px #0000001a;--tw-shadow-colored:0 1px 3px 0 var(--tw-shadow-color),0 1px 2px -1px var(--tw-shadow-color)}.shadow,.shadow-lg{box-shadow:var(--tw-ring-offset-shadow,0 0 #0000),var(--tw-ring-shadow,0 0 #0000),var(--tw-shadow)}.shadow-lg{--tw-shadow:0 10px 15px -3px #0000001a,0 4px 6px -4px #0000001a;--tw-shadow-colored:0 10px 15px -3px var(--tw-shadow-color),0 4px 6px -4px var(--tw-shadow-color)}.shadow-md{--tw-shadow:0 4px 6px -1px #0000001a,0 2px 4px -2px #0000001a;--tw-shadow-colored:0 4px 6px -1px var(--tw-shadow-color),0 2px 4px -2px var(--tw-shadow-color);box-shadow:var(--tw-ring-offset-shadow,0 0 #0000),var(--tw-ring-shadow,0 0 #0000),var(--tw-shadow)}.outline-none{outline:2px solid #0000;outline-offset:2px}.outline{outline-style:solid}.outline-dashed{outline-style:dashed}.outline-dotted{outline-style:dotted}.outline-double{outline-style:double}.ring{--tw-ring-offset-shadow:var(--tw-ring-inset) 0 0 0 var(--tw-ring-offset-width) var(--tw-ring-offset-color);--tw-ring-shadow:var(--tw-ring-inset) 0 0 0 calc(3px + var(--tw-ring-offset-width)) var(--tw-ring-color);box-shadow:var(--tw-ring-offset-shadow),var(--tw-ring-shadow),var(--tw-shadow,0 0 #0000)}.ring-inset{--tw-ring-inset:inset}.blur{--tw-blur:blur(8px)}.blur,.drop-shadow{filter:var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)}.drop-shadow{--tw-drop-shadow:drop-shadow(0 1px 2px #0000001a) drop-shadow(0 1px 1px #0000000f)}.grayscale{--tw-grayscale:grayscale(100%)}.grayscale,.invert{filter:var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)}.invert{--tw-invert:invert(100%)}.sepia{--tw-sepia:sepia(100%);filter:var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)}.\!filter{filter:var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)!important}.filter{filter:var(--tw-blur) var(--tw-brightness) var(--tw-contrast) var(--tw-grayscale) var(--tw-hue-rotate) var(--tw-invert) var(--tw-saturate) var(--tw-sepia) var(--tw-drop-shadow)}.filter-none{filter:none}.backdrop-blur{--tw-backdrop-blur:blur(8px)}.backdrop-blur,.backdrop-grayscale{-webkit-backdrop-filter:var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia);backdrop-filter:var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)}.backdrop-grayscale{--tw-backdrop-grayscale:grayscale(100%)}.backdrop-invert{--tw-backdrop-invert:invert(100%)}.backdrop-invert,.backdrop-sepia{-webkit-backdrop-filter:var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia);backdrop-filter:var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)}.backdrop-sepia{--tw-backdrop-sepia:sepia(100%)}.backdrop-filter{-webkit-backdrop-filter:var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia);backdrop-filter:var(--tw-backdrop-blur) var(--tw-backdrop-brightness) var(--tw-backdrop-contrast) var(--tw-backdrop-grayscale) var(--tw-backdrop-hue-rotate) var(--tw-backdrop-invert) var(--tw-backdrop-opacity) var(--tw-backdrop-saturate) var(--tw-backdrop-sepia)}.backdrop-filter-none{-webkit-backdrop-filter:none;backdrop-filter:none}.transition{transition-property:color,background-color,border-color,text-decoration-color,fill,stroke,opacity,box-shadow,transform,filter,-webkit-backdrop-filter;transition-property:color,background-color,border-color,text-decoration-color,fill,stroke,opacity,box-shadow,transform,filter,backdrop-filter;transition-property:color,background-color,border-color,text-decoration-color,fill,stroke,opacity,box-shadow,transform,filter,backdrop-filter,-webkit-backdrop-filter;transition-timing-function:cubic-bezier(.4,0,.2,1);transition-duration:.15s}.ease-in{transition-timing-function:cubic-bezier(.4,0,1,1)}.ease-in-out{transition-timing-function:cubic-bezier(.4,0,.2,1)}.ease-out{transition-timing-function:cubic-bezier(0,0,.2,1)}.contain-none{contain:none}.contain-content{contain:content}.contain-strict{contain:strict}.contain-size{--tw-contain-size:size}.contain-inline-size,.contain-size{contain:var(--tw-contain-size) var(--tw-contain-layout) var(--tw-contain-paint) var(--tw-contain-style)}.contain-inline-size{--tw-contain-size:inline-size}.contain-layout{--tw-contain-layout:layout}.contain-layout,.contain-paint{contain:var(--tw-contain-size) var(--tw-contain-layout) var(--tw-contain-paint) var(--tw-contain-style)}.contain-paint{--tw-contain-paint:paint}.contain-style{--tw-contain-style:style;contain:var(--tw-contain-size) var(--tw-contain-layout) var(--tw-contain-paint) var(--tw-contain-style)}.content-\[\'this-is-also-valid\]-weirdly-enough\'\]{--tw-content:"this-is-also-valid]-weirdly-enough";content:var(--tw-content)}.forced-color-adjust-auto{forced-color-adjust:auto}.forced-color-adjust-none{forced-color-adjust:none}@media (min-width:640px){.sm\:container{width:100%}@media (min-width:640px){.sm\:container{max-width:640px}}@media (min-width:768px){.sm\:container{max-width:768px}}@media (min-width:1024px){.sm\:container{max-width:1024px}}@media (min-width:1280px){.sm\:container{max-width:1280px}}@media (min-width:1536px){.sm\:container{max-width:1536px}}}.hover\:bg-slate-950:hover{--tw-bg-opacity:1;background-color:rgb(2 6 23/var(--tw-bg-opacity))}.hover\:font-bold:hover{font-weight:700}.before\:hover\:text-center:hover:before,.hover\:before\:text-center:hover:before{content:var(--tw-content);text-align:center}.focus\:hover\:text-center:hover:focus,.hover\:focus\:text-center:focus:hover{text-align:center}@media (min-width:640px){.sm\:underline{text-decoration-line:underline}}@media (prefers-color-scheme:dark){@media (min-width:1024px){.dark\:lg\:hover\:\[paint-order\:markers\]:hover{paint-order:markers}}}
	</style>
	
	<body class="font-mono">
		<header class="w-full bg-slate-500 p-4 flex items-center pt-4 shadow-md">
		<span class="font-bold font">V{{.Version}}</span>
			<div class="flex  justify-center flex-1">
			
				<h1 class="font-semibold text-3xl">
					DB Compare
				</h1>
			</div>
			<div>
			<div>
            <a href="https://github.com/LuanChagas" target="_blank" class="underline">By Luan Chagas</a>
        </div>
			</div>
		</header>
		<main>
		<section class="w-full flex flex-col items-center p-7 border">
		<div class="w-1/2 border  p-6 rounded-lg flex flex-col">
			<h1 class="text-xl font-bold">Stats</h1>
			<span>Tempo Geral: {{.Stats.TempoGeral}}</span>
			<span>Tempo processamento de dados Primários: {{.Stats.TempoProcessamentoDadosPrimarios}} </span>
			<span>Tempo processamento de dados secundario: {{.Stats.TempoProcessamentoDadosSecundarios}} </span>
			<span>Tempo Comparação: {{.Stats.TempoComparacao}} </span>
			<span>Tamanho de dados Primarios: {{.Stats.TamanhoDadosPrimarios}} </span>
			<span>Tamanho de dados Secundários: {{.Stats.TamanhoDadosSecundarios}} </span>
		</div>
	   
	   
	</section>

			<section class="w-full flex flex-col items-center p-7 ">
				<div class="w-10/12 border shadow-lg p-6 rounded-lg">
					<h3 class="pb-2 text-2xl border-b-2 border-b-black">Tabelas</h3>
					<div class="flex justify-center ">
						<table class="min-w-full bg-white rounded-lg border-separate border-spacing-3">
							<caption class="caption-bottom">
								Tabelas existentes nos Bancos de dados
							</caption>
							<thead>
								<tr class="">
									<th class="py-2 px-4 rounded-md"></th>
									<th class="py-2 px-4 bg-gray-200 rounded-md">{{.BancoPrimario}}</th>
									<th class="py-2 px-4 bg-gray-200 rounded-md">{{.BancoSecundario}}</th>
								</tr>
							</thead>
							<tbody>





								{{range  $nome, $existe := .Resultado.Tabelas.Tabela}}
								<tr>
									<td class="border px-4 py-2 rounded-md ">{{$nome}}</td>
									<td class="border px-4 py-2 text-center rounded-md 
									{{if $existe.Primario}}bg-green-200{{else}}bg-red-200{{end}}
									">{{if $existe.Primario}}Sim{{else}}Não{{end}}</td>
									<td class="border px-4 py-2 text-center rounded-md
									{{if $existe.Secundario}}bg-green-200 {{else}}bg-red-200{{end}}
									">{{if $existe.Secundario}}Sim{{else}}Não{{end}}</td>
								</tr>
								{{end}}






							</tbody>
						</table>
					</div>
				</div>
			</section>
			<section class="w-full flex justify-center">
				<button id="btnPrimario"
					class="bg-slate-800 p-4 rounded-md shadow-lg hover:bg-slate-950 ">
					<span class="text-green-500">{{.BancoPrimario}} -></span> <span class="text-white"> Compare </span> <span
						class="text-red-600">-> {{.BancoSecundario}} </span>
				</button>
				<button id="btnSecundario" class="bg-slate-800 p-4 rounded-md shadow-lg hover:bg-slate-950 hidden">
					<span class="text-green-500">{{.BancoSecundario}} -></span> <span class="text-white"> Compare </span> <span
						class="text-red-600">-> {{.BancoPrimario}} </span>
				</button>
			</section>
			<section id="primario" class="w-full flex flex-col items-center p-7 ">
			
				<section class="w-10/12 border shadow-lg p-6 rounded-lg">
					<h3 class="pb- text-2xl border-b-2 border-b-black">Campos</h3>
					<section class="flex flex-col gap-10 pt-10">


				{{range $tabela := .Resultado.BancoPrimario }}
					{{if and  $tabela.Diferenca $tabela.TabelaExiste }}


				<section class="flex flex-col border-2 rounded-lg shadow-lg p-6">
				<div class=" ml-5 text-xl border-b border-b-slate-400">
				<h3 class=" text-xl ">{{$tabela.Tabela}}</h3>
			</div>
			<div class="flex justify-center">
			<table class="min-w-full bg-white rounded-lg border-separate border-spacing-3">
				
				<thead>
					<tr class="">
						<th class="py-2 px-4 rounded-md"></th>
						<th class="py-2 px-4 bg-gray-200 rounded-md">Comparaçao</th>
						<th class="py-2 px-4 bg-gray-200 rounded-md">{{$.BancoPrimario}}</th>
						<th class="py-2 px-4 bg-gray-200 rounded-md">{{$.BancoSecundario}}</th>
					</tr>
				</thead>
				<tbody>


				{{range $dadosCampos := $tabela.Colunas}}


					<tr>
						<td class="border px-4 py-2 rounded-md ">{{$dadosCampos.Nome}}</td>
						<td class="border px-4 py-2 text-center rounded-md  ">{{$dadosCampos.TipoComparacao}}</td>
						<td class="border px-4 py-2 text-center rounded-md  ">{{$dadosCampos.Primario}}</td>
						<td class="border px-4 py-2 text-center rounded-md ">{{$dadosCampos.Secundario}}</td>
					</tr>


				{{end}}

				{{range $dadosCampos := $tabela.Chaves}}
					<tr>
						<td class="border px-4 py-2 rounded-md ">{{$dadosCampos.Nome}}</td>
						<td class="border px-4 py-2 text-center rounded-md  ">{{$dadosCampos.TipoComparacao}}</td>
						<td class="border px-4 py-2 text-center rounded-md  ">{{$dadosCampos.Primario}}</td>
						<td class="border px-4 py-2 text-center rounded-md ">{{$dadosCampos.Secundario}}</td>
					</tr>
				{{end}}
				</tbody>
			</table>
		</div>
	</section>
	{{end}}
				{{end}}
</section>
			</section>
			</section>
			<section id="secundario"
				class="w-full  flex-col items-center p-7 hidden ">
				<section class="w-10/12 border shadow-lg p-6 rounded-lg">
					<h3 class="pb- text-2xl border-b-2 border-b-black">Campos </h3>
					<section class="flex flex-col gap-10 pt-10">	
					
				{{range $tabela := .Resultado.BancoSecundario }}
				{{if and  $tabela.Diferenca $tabela.TabelaExiste }}


			<section class="flex flex-col border-2 rounded-lg shadow-lg p-6">
			<div class=" ml-5 text-xl border-b border-b-slate-400">
			<h3 class=" text-xl ">{{$tabela.Tabela}}</h3>
		</div>
		<div class="flex justify-center">
		<table class="min-w-full bg-white rounded-lg border-separate border-spacing-3">
			
			<thead>
				<tr class="">
					<th class="py-2 px-4 rounded-md"></th>
					<th class="py-2 px-4 bg-gray-200 rounded-md">Comparaçao</th>
					<th class="py-2 px-4 bg-gray-200 rounded-md">{{$.BancoSecundario}}</th>
					<th class="py-2 px-4 bg-gray-200 rounded-md">{{$.BancoPrimario}}</th>
				</tr>
			</thead>
			<tbody>


			{{range $dadosCampos := $tabela.Colunas}}


				<tr>
					<td class="border px-4 py-2 rounded-md ">{{$dadosCampos.Nome}}</td>
					<td class="border px-4 py-2 text-center rounded-md  ">{{$dadosCampos.TipoComparacao}}</td>
					<td class="border px-4 py-2 text-center rounded-md  ">{{$dadosCampos.Primario}}</td>
					<td class="border px-4 py-2 text-center rounded-md ">{{$dadosCampos.Secundario}}</td>
				</tr>


			{{end}}

			{{range $dadosCampos := $tabela.Chaves}}
				<tr>
					<td class="border px-4 py-2 rounded-md ">{{$dadosCampos.Nome}}</td>
					<td class="border px-4 py-2 text-center rounded-md  ">{{$dadosCampos.TipoComparacao}}</td>
					<td class="border px-4 py-2 text-center rounded-md  ">{{$dadosCampos.Primario}}</td>
					<td class="border px-4 py-2 text-center rounded-md ">{{$dadosCampos.Secundario}}</td>
				</tr>
			{{end}}
			</tbody>
		</table>
	</div>
</section>


{{end}}
			{{end}}	
					</section>
				</section>
			</section>
		</main>
		<script>
			btnPrimario.onclick = () => {
				btnPrimario.classList.add('hidden')
				primario.classList.add('hidden')
				btnSecundario.classList.remove('hidden')
				secundario.classList.remove('hidden')
				secundario.classList.add('flex')
			}
			btnSecundario.onclick = () => {
				btnSecundario.classList.add('hidden')
				secundario.classList.add('hidden')
				btnPrimario.classList.remove('hidden')
				primario.classList.remove('hidden')
				primario.classList.add('flex')
			}
		</script>
	</body>
	</html>`

	temp := template.New("resultado")
	temp, err := temp.Parse(strTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create("resultado.html")
	if err != nil {
		return err
	}

	defer file.Close()

	err = temp.Execute(file, dadosOutput)
	if err != nil {
		return err
	}

	color.Green("Arquivo resultado.html gerado na raiz do executável")
	return nil
}
