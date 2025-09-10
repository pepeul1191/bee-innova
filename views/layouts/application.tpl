<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.PageTitle}} | TicketMaster</title>
  <link rel="icon" href="/static/icon.png" type="image/png">
  <link rel="apple-touch-icon" href="/static/icon.png">
  
  <!-- Bootstrap CSS -->
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet">
  <!-- Font Awesome -->
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css">
  
  {{ if .Styles }}
    {{ range .Styles }}
      <link rel="stylesheet" href="{{.}}">
    {{ end }}
  {{ end }}
  
  <!-- Estilos personalizados -->
  <style>
    .sidebar {
      width: 250px;
      height: 100vh;
      position: fixed;
      top: 0;
      left: 0;
      z-index: 1000;
      transition: all 0.3s;
    }
    
    .main-content {
      margin-left: 250px;
      transition: all 0.3s;
      padding-top: 70px;
      min-height: 100vh;
    }
    
    .sidebar.collapsed {
      margin-left: -250px;
    }
    
    .main-content.expanded {
      margin-left: 0;
    }
    
    .nav-link {
      color: rgba(255, 255, 255, 0.8) !important;
      transition: all 0.3s;
    }
    
    .nav-link:hover, .nav-link.active {
      color: #fff !important;
      background-color: rgba(255, 255, 255, 0.1);
    }
    
    .navbar {
      z-index: 999;
    }
    
    #userDropdown span.user-name {
      margin-left: 0.5rem;
    }

    .return-nav{
      text-decoration: none !important;
    }

    .return-nav i{
      margin-right: 14px !important;
    }

    .btn-row{
      padding-top: 32px;
    }

    .main-content .py-4{
      padding-bottom: 0rem !important; 
    }

    .sidebar.collapsed {
      transform: translateX(-250px);
    }

    .content-wrapper {
      margin-left: 250px;
      padding: 20px;
      transition: all 0.3s;
    }

    .content-wrapper .py-4{
      padding-top: 0rem !important;
    }

    .sidebar.collapsed ~ .content-wrapper {
      margin-left: 0;
    }

    @media (max-width: 768px) {
      .sidebar {
        transform: translateX(-250px);
      }
      
      .sidebar.show {
        transform: translateX(0);
      }
      
      .content-wrapper {
        margin-left: 0;
      }
    }

    /* En tu archivo dashboard.css o styles.css */
    .sidebar .nav-link.active {
      background-color: white !important;
      color: black !important;
      font-weight: bold;
    }

    /* Opcional: Efecto hover */
    .sidebar .nav-link:hover:not(.active) {
      background-color: rgba(255, 255, 255, 0.1);
    }

    .nav-link{
      color: #FFF;
    }

    .nav-link:hover{
      background-color: #dadada !important;
      color: #343434;
    }

    /* formularios */
    h2 a {
      text-decoration: none;
    }
    .suggestion-list {
      position: absolute;
      z-index: 1000;
      background: #fff;
      width: inherit;
      max-height: 300px;
      overflow-y: auto;
      margin-top: 2px;
      border-radius: 4px;
    }

    .suggestion-item {
      padding: 8px 12px;
      cursor: pointer;
      border: 1px solid #eee;
    }

    .suggestion-item:hover, 
    .suggestion-item.active {
      background-color: #e1e1e1;
      border: 0px;
    }
    .spinner {
      animation: spin 1s linear infinite;
    }
    @keyframes spin {
      0% { transform: rotate(0deg); }
      100% { transform: rotate(360deg); }
    }
  </style>
</head>
<body class="dashboard-layout">
  <!-- Incluir Navbar -->
  {{template "partials/_navbar.tpl" .}}
  
  <!-- Incluir Sidebar -->
  {{template "partials/_sidebar.tpl" .}}

  <!-- Contenido principal -->
  <div class="main-content" id="mainContent">
    <div class="container-fluid">
      {{.LayoutContent}}
    </div>
  </div>

  <!-- Bootstrap JS -->
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js"></script>
  
  {{ if .Scripts }}
    {{ range .Scripts }}
      <script src="{{.}}"></script>
    {{ end }}
  {{ end }}

  <!-- Script para manejar la sesión del usuario -->
  <script>
    document.addEventListener('DOMContentLoaded', function() {
      // Función para obtener datos de sesión
      async function getSessionData() {
        try {
          const response = await fetch('/api/v1/session', {
            method: 'GET',
            credentials: 'include'
          });
          
          if (!response.ok) {
            throw new Error('Error al obtener sesión');
          }
          
          const data = await response.json();
          return data;
        } catch (error) {
          console.error('Error:', error);
          return { success: false };
        }
      }

      // Función para actualizar la UI con los datos de sesión
      async function updateUIBasedOnSession() {
        const sessionData = await getSessionData();
        const userNameElement = document.getElementById('userName');
        const authSection = document.getElementById('authSection');
        
        if (sessionData.success && sessionData.data && sessionData.data.user) {
          const user = sessionData.data.user;
          userNameElement.textContent = user.full_name || user.username || 'Usuario';
          
          authSection.innerHTML = `
            <a class="dropdown-item" href="/sign-out">
              <i class="fas fa-sign-out-alt me-2"></i> Cerrar sesión
            </a>
          `;
        } else {
          userNameElement.textContent = 'Invitado';
          authSection.innerHTML = `
            <a class="dropdown-item" href="/sign-in">
              <i class="fas fa-sign-in-alt me-2"></i> Iniciar sesión
            </a>
          `;
        }
      }

      // Toggle sidebar
      document.getElementById('sidebarToggle').addEventListener('click', function() {
        const sidebar = document.querySelector('.sidebar');
        const mainContent = document.getElementById('mainContent');
        
        sidebar.classList.toggle('collapsed');
        mainContent.classList.toggle('expanded');
      });

      // Actualizar UI con datos de sesión
      updateUIBasedOnSession();
      
      // Opcional: Actualizar cada cierto tiempo
      setInterval(updateUIBasedOnSession, 60000);
    });
  </script>
</body>
</html>