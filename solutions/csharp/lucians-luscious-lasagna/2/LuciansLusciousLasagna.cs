class Lasagna
{
    public const int expectedMinutes = 40;
    // Expression-bodied method
    public int ExpectedMinutesInOven() => expectedMinutes;
    // TODO: define the 'RemainingMinutesInOven()' method
    public int RemainingMinutesInOven(int inOven) {
        return expectedMinutes - inOven;
    }
    // TODO: define the 'PreparationTimeInMinutes()' method
    public int PreparationTimeInMinutes(int layers) {
        const int minutesPerLayer = 2;
        return layers * 2;
    }
    // TODO: define the 'ElapsedTimeInMinutes()' method
    public int ElapsedTimeInMinutes(int layers, int inOven) {
        return PreparationTimeInMinutes(layers) + inOven;
    }
}
