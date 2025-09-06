<!-- Slider -->
<div id="mainCarousel" class="carousel slide" data-bs-ride="carousel">
  <div class="carousel-inner">
    <div class="carousel-item active">
      <img src="{{ GetEnv "STATIC_URL" }}common/img/slider.webp" class="d-block w-100" alt="Emprendimiento">
      <div class="carousel-caption d-none d-md-block">
        <h2>Impulsamos tu idea de negocio <strong>{{ GetEnv "DB_PORT" }}</strong></h2>
        <p>Apoyo, mentoría y recursos para emprendedores.</p>
      </div>
    </div>
    <div class="carousel-item">
      <img src="{{ GetEnv "STATIC_URL" }}common/img/slider.webp" class="d-block w-100" alt="Reunión">
      <div class="carousel-caption d-none d-md-block">
        <h2>Conecta con mentores y expertos</h2>
        <p>Forma parte de nuestra red de apoyo empresarial.</p>
      </div>
    </div>
    <div class="carousel-item">
      <img src="{{ GetEnv "STATIC_URL" }}common/img/slider.webp" class="d-block w-100" alt="Tecnología">
      <div class="carousel-caption d-none d-md-block">
        <h2>Espacios modernos y colaborativos</h2>
        <p>Accede a coworking y salas de innovación.</p>
      </div>
    </div>
  </div>
  <button class="carousel-control-prev" type="button" data-bs-target="#mainCarousel" data-bs-slide="prev">
    <span class="carousel-control-prev-icon"></span>
  </button>
  <button class="carousel-control-next" type="button" data-bs-target="#mainCarousel" data-bs-slide="next">
    <span class="carousel-control-next-icon"></span>
  </button>
</div>

<!-- Sobre Nosotros -->
<section id="about" class="py-5 bg-light">
  <div class="container text-center">
    <h2 class="mb-4">¿Quiénes Somos?</h2>
    <p class="lead">Somos un espacio dedicado a potenciar el talento emprendedor. Ofrecemos formación, acompañamiento y una red de contactos para transformar ideas en empresas sostenibles.</p>
  </div>
</section>

<!-- Programas / Servicios -->
<section id="services" class="py-5 services">
  <div class="container">
    <h2 class="text-center mb-4">Nuestros Programas</h2>
    <div class="row">
      <div class="col-md-4">
        <div class="card h-100 text-center">
          <div class="card-body">
            <h5 class="card-title">Pre-incubación</h5>
            <p class="card-text">Te ayudamos a validar tu idea de negocio y estructurar tu modelo.</p>
          </div>
        </div>
      </div>
      <div class="col-md-4">
        <div class="card h-100 text-center">
          <div class="card-body">
            <h5 class="card-title">Incubación</h5>
            <p class="card-text">Acompañamiento completo para lanzar tu startup al mercado.</p>
          </div>
        </div>
      </div>
      <div class="col-md-4">
        <div class="card h-100 text-center">
          <div class="card-body">
            <h5 class="card-title">Aceleración</h5>
            <p class="card-text">Potenciamos negocios en crecimiento con mentoría, inversión y escalabilidad.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>