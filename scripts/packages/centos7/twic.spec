Name: twic
Version: %{_version}
Release: %{_release}%{?dist}
Summary: HBM TWIC
Group: Tools/Docker

License: GPL

URL: https://github.com/kassisol/twic
Vendor: Kassisol
Packager: Kassisol <support@kassisol.com>

BuildArch: x86_64
BuildRoot: %{_tmppath}/%{name}-buildroot

Source: twic.tar.gz

%description
HBM TWIC is an open source project for managing Docker certificates to connect to the Docker daemon using TLS.

%prep
%setup -n %{name}

%install
# install binary
install -d $RPM_BUILD_ROOT/%{_bindir}
install -p -m 755 twic $RPM_BUILD_ROOT/%{_bindir}/

# add bash completions
install -d $RPM_BUILD_ROOT/usr/share/bash-completion/completions
install -p -m 644 shellcompletion/bash $RPM_BUILD_ROOT/usr/share/bash-completion/completions/twic

# install manpages
install -d $RPM_BUILD_ROOT/%{_mandir}/man8
install -p -m 644 man/man8/*.8 $RPM_BUILD_ROOT/%{_mandir}/man8

# list files owned by the package here
%files
#%doc README.md
%{_bindir}/twic
/usr/share/bash-completion/completions/twic
%doc
/%{_mandir}/man8/*

%postun
rm -f %{_bindir}/twic

%clean
rm -rf $RPM_BUILD_ROOT
