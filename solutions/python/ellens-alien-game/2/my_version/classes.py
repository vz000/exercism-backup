"""Solution to Ellen's Alien Game exercise."""

"""
In this solution is_alive is treated as an attribute and not a method. 
It takes care of cases where the value could be changed, since the attribute itself can be
accessed freely.
Using properties prevents _is_alive to be changed to True when, for example,
health is already 0.

NOTE: Pylint may return:
    [C0104 disallowed-name] : ["Disallowed name "y"] was reported by Pylint.
"""

class Coordinates():
    def __set_name__(self, owner, coord):
        self.coord = coord

    def __get__(self, obj, objtype=None):
        return obj.__dict__.get(self.coord)

    def __set__(self, obj, value):
        if isinstance(value, int):
            obj.__dict__[self.coord] = value
        else:
            raise ValueError("Coordinate value is not valid.")
    
class Alien:
    """Create an Alien object with location x_coordinate and y_coordinate.

    Attributes
    ----------
    (class)total_aliens_created: int
    x_coordinate: int - Position on the x-axis.
    y_coordinate: int - Position on the y-axis.
    health: int - Number of health points.

    Methods
    -------
    hit(): Decrement Alien health by one point.
    is_alive(): Return a boolean for if Alien is alive (if health is > 0).
    teleport(new_x_coordinate, new_y_coordinate): Move Alien object to new coordinates.
    collision_detection(other): Implementation TBD.
    """

    x_coordinate = Coordinates()
    y_coordinate = Coordinates()
    alien_counter = 0

    def __init__(self, x, y):
        self.x_coordinate = x
        self.y_coordinate = y
        self.health = 3
        self._is_alive = True
        Alien.alien_counter += 1

    @classmethod
    def total_aliens_created(cls):
        return cls.alien_counter

    def hit(self):
        if (self.health - 1) <= 0:
            self.health = 0
            self._is_alive = False
        else:
            self.health -= 1

    @property
    def is_alive(self):
        return self._is_alive

    @is_alive.setter
    def is_alive(self, value):
        if self.health > 0 and value == False:
            raise ValueError("Health is not 0.")
        elif self.health == 0 and value == True:
            raise ValueError("Health is 0 and alien cannot be alive.")
        self._is_alive = value

    def teleport(self, x, y):
        self.x_coordinate = x
        self.y_coordinate = y

    def collision_detection(self, other_alien):
        pass

def new_aliens_collection(coordinates: list) -> list:
    aliens_list = list()
    for new_coords in coordinates:
        aliens_list.append(Alien(*new_coords))
    
    return aliens_list
