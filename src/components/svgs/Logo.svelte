<script>
	export let size = 24;
	export let color = '#000000'; // Color por defecto
  export let styles = ''; // Color por defecto

	$: sizePx = `${size}px`;
	$: fillColor = color; // Usar la variable de color para el atributo fill
</script>

<svg style="{styles}" width="{size}mm" height="{size}mm" version="1.1" viewBox="0 0 54 54" xmlns="http://www.w3.org/2000/svg">
  <g transform="translate(0,-243)">
    <g transform="matrix(.82731 0 0 .82731 -62.827 133.8)">
      <path d="m87.619 144.65 12.804 1.8899 1.8899 12.709-12.946-1.7954z" fill="{fillColor}"/>
      <path d="m114.2 171.22 12.804 1.8899 1.8899 12.709-12.946-1.7954z" fill="{fillColor}"/>
      <path d="m108.16 136.19 7.7945 10.332-7.5736 10.38-7.9618-10.365z" fill="{fillColor}"/>
      <path d="m108.16 173.69 7.7945 10.332-7.5736 10.38-7.9618-10.365z" fill="{fillColor}"/>
      <path d="m128.76 144.78-1.8899 12.804-12.709 1.8899 1.7954-12.946z" fill="{fillColor}"/>
      <path d="m102.31 171.23-1.8899 12.804-12.709 1.8899 1.7954-12.946z" fill="{fillColor}"/>
      <path d="m116.68 165.34 10.353-7.7677 10.36 7.6004-10.385 7.9349z" fill="{fillColor}"/>
      <path d="m79.167 165.34 10.353-7.7677 10.36 7.6004-10.385 7.9349z" fill="{fillColor}"/>
    </g>
  </g>
</svg>