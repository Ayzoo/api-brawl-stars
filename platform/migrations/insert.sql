INSERT INTO brawlers_gadget (brawler_id, primary_gadget, second_gadget) 
VALUES (1, 'FAST FORWARD', 'CLAY PIGEONS');


INSERT INTO brawlers_stellar (brawler_id, primary_stellar, second_stellar)
VALUES (1, 'SHELL SHOCK', 'BAND-AID');

INSERT INTO super (brawler_id, name, description, damage, scope)
VALUES (1, 'SUPER SHELL', 'Shelly Super Shell obliterates both cover and enemies. Any survivors get knocked back.', '11 x 480', 'LONG');

INSERT INTO brawlers (name, image, type, category, stellar, gadget, super) 
VALUES ("Shelly", "https://i.postimg.cc/hjcswwKD/Shelly.png", "Starting Brawler", "Damage Dealer", 1, 1, 1);