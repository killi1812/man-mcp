Name:           man-mcp
Version:        0.0.1
Release:        1%{?dist}
Summary:        Model Context Protocol server for local system man pages

License:        GPL-3.0-or-later
URL:            https://github.com/killi1812/man-mcp
Source0:        %{url}/archive/v%{version}/%{name}-%{version}.tar.gz

BuildRequires:  golang >= 1.24
Requires:       man-db

%description
man-mcp is a Model Context Protocol (MCP) server that provides AI coding
assistants with direct access to local system manual pages as a reliable
source of truth.

%prep
%autosetup -p1

%build
cd src
go build -ldflags="-X 'github.com/killi1812/man-mcp/app.Build=prod' -X 'github.com/killi1812/man-mcp/app.Version=%{version}'" -o %{name} .

%install
install -Dpm 0755 src/%{name} %{buildroot}%{_bindir}/%{name}

%files
%license LICENSE
%doc README.md
%{_bindir}/%{name}

%changelog
* Sun Sep 29 2026 Fran Cvok <cvok.fran@gmail.com> - 0.0.1-1
- Initial RPM release
