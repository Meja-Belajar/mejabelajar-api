import { registerService } from '../src/apis/services/userService';

// Mocking the fetch function
global.fetch = jest.fn();

describe('registerService', () => {
  afterEach(() => {
    jest.clearAllMocks();
  });

  it('should register user successfully', async () => {
    const mockResponseData = {
      code: 200,
      message: 'Success: Account has been created',
      data: {  }
    };

    // Mock successful fetch response
    const mockResponse = {
      ok: true,
      json: jest.fn().mockResolvedValue(mockResponseData),
      headers: new Headers(),
      redirected: false,
      status: 200,
      statusText: 'OK',
      type: 'basic',
      url: 'http://example.com'
    };

    (fetch as jest.MockedFunction<typeof fetch>).mockResolvedValue(mockResponse);

    const requestData = {
      user_name: 'testuser',
      email: 'test@example.com',
      password: 'password',
      phone_number: '1234567890',
      bod: '2000-01-01', 
      confirm_password: 'password',
      created_by: 'admin'
    };

    const result = await registerService(requestData);
    expect(result).toEqual(mockResponseData);
  });

  it('should throw an error if registration fails', async () => {
    // Mock failed fetch response
    (fetch as jest.MockedFunction<typeof fetch>).mockResolvedValue({
      ok: false,
      statusText: 'Internal Server Error'
    });

    const requestData = {
      user_name: 'testuser',
      email: 'test@example.com',
      password: 'password',
      phone_number: '1234567890',
      bod: '2000-01-01', // Example date of birth
      confirm_password: 'password',
      created_by: 'admin'
    };

    await expect(registerService(requestData)).rejects.toThrow('Failed to register user');
  });
});
