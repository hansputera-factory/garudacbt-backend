{pkgs, ...}: {
    packages = [
        pkgs.go_1_22
        pkgs.nodejs-slim_21
    ];

    # Enabling MySQL
    services.mysql.enable = true;
    services.mysql.package = pkgs.mysql80;

    # Enabling Redis
    services.redis.enable = true;
    services.redis.port = 6335;
}
