<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.PageTitle}} | TicketMaster</title>
  <link rel="icon" href="/static/icon.png" type="image/png">
  <link rel="apple-touch-icon" href="/static/icon.png">  
  {{ if .Styles }}
    {{ styles .Styles }}
  {{ end }}
  
  <!-- Estilos personalizados -->
  <style>
 
  </style>
</head>
<body class="dashboard-layout">
  <div class="wrapper">
    <!-- Incluir Navbar -->
    {{template "partials/_navbar.tpl" .}}
    
    <!-- Incluir Sidebar -->
    {{template "partials/_sidebar.tpl" .}}

    <!-- Contenido principal -->
    <div class="content-wrapper">
      <main class="main-content">
        <div class="container-fluid">
          {{.LayoutContent}}
        </div>
      </main>
    </div>  
  </div>

  {{ if .Scripts }}
    {{ scripts .Scripts }}
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